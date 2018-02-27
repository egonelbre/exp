package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/websocket"

	"github.com/loov/watchrun/watch"
)

var WebsocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeReloadJS(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Upgrade") != "" {
		conn, err := WebsocketUpgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		ServeReloadChanges(conn)
		return
	}

	url := "ws://" + r.Host + r.RequestURI
	data := strings.Replace(ReloaderJS, "DEFAULT_HOST", url, -1)
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(data))
}

type Message struct {
	Type string  `json:"type"`
	Data Changes `json:"data"`
}

type Changes []Change
type Change struct {
	Kind     string    `json:"kind"`
	Path     string    `json:"path"`
	Modified time.Time `json:"modified"`
}

func ServeReloadChanges(conn *websocket.Conn) {
	defer conn.Close()

	fmt.Println("CONNECTED", conn.LocalAddr())
	defer fmt.Println("DISCONNECTED", conn.LocalAddr())

	watcher := watch.New(*interval, nil, nil, []string{"*.js", "*.css"}, true)
	defer watcher.Stop()

	go func() {
		defer conn.Close()
		defer watcher.Stop()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Websocket Error: %v", err)
				}
				break
			}
		}
	}()

	for changeset := range watcher.Changes {
		message := Message{
			Type: "changes",
		}

		for _, change := range changeset {
			rel, err := filepath.Rel(ClientRootDir, change.Path)
			if err != nil {
				rel = change.Path
			}
			path := filepath.ToSlash(rel)
			if path[0] != '/' {
				path = "/" + path
			}
			message.Data = append(message.Data, Change{
				Kind:     change.Kind,
				Path:     path,
				Modified: change.Modified,
			})
		}

		if err := conn.WriteJSON(message); err != nil {
			return
		}
	}
}

const ReloaderJS = `
(function(global){
	"use strict";

	global.Reloader = Reloader;
	global.Reloader.defaultHost = "DEFAULT_HOST";
	function Reloader(host) {
		var reloader = this;
		this.host = host || Reloader.defaultHost;

		this.socket = new WebSocket(this.host);
		this.socket.onopen = function(ev) {
			console.info("RELOADER: socket openened ", ev);
		};
		this.socket.onerror = function(ev) {
			console.info("RELOADER: socket encountered error ", ev);
		};
		this.socket.onclose = function(ev) {
			console.info("RELOADER: socket closed ", ev);
			reloader.ondisconnected();
		};
		this.socket.onmessage = this.message.bind(this);
		this.changeset = 0;
	}

	Reloader.prototype = {
		message: function(ev) {
			var reloader = this;

			reloader.changeset++;
			if (reloader.changeset <= 1) {
				return;
			}

			var message = JSON.parse(ev.data);
			reloader["on" + message.type].call(reloader, message.data);
		},
		ondisconnected: function(){
			console.info("RELOADER: waiting server restart");
			this._tryreconnect();
		},
		_tryreconnect: function(){
			var reloader =  this;
			window.setTimeout(function(){
				var probe = new WebSocket(reloader.host);
				probe.onopen = function(ev){
					window.location.reload();
				};
				probe.onerror = function(){ reloader._tryreconnect(); };
			}, 500);
		},
		onchanges: function(changes) {
			if(this.onbeforechanges){
				if(this.onbeforechanges(changes)){
					return;
				}
			}

			function pathext(name) {
				var i = name.lastIndexOf(".");
				if (i < 0) {
					return "";
				}
				return name.substring(i);
			}

			function makeasset(name, modified) {
				switch (pathext(name)) {
					case ".js":
						var asset = document.createElement("script");
						asset.id = name;
						asset.src = name + "?" + modified;
						return asset;
					case ".css":
						var asset = document.createElement("link");
						asset.id = name;
						asset.href = name + "?" + modified;
						asset.rel = "stylesheet";
						return asset;
				}
				return undefined;
			}

			function findasset(name) {
				var el = document.getElementById(name);
				if (el) {
					return el;
				}

				switch (pathext(name)) {
					case ".js":
						return document.querySelector("script[src='" + name + "']");
					case ".css":
						return document.querySelector("link[href='" + name + "']");
				}
				return undefined;
			}

			function inject(change) {
				var el = findasset(change.path);
				if (el) {
					el.remove();
				}

				var asset = makeasset(change.path, change.modified);
				if (asset) {
					document.head.appendChild(asset);
				} else {
					location.reload();
				}
			}

			function remove(change) {
				var el = findasset(change.path);
				if (el) {
					el.remove();
				}
			}

			function modify(change) {
				remove(change);
				inject(change);
			}

			for (var i = 0; i < changes.length; i++) {
				var change = changes[i];
				switch (change.kind) {
					case "create":
						inject(change);
						break;
					case "delete":
						remove(change);
						break;
					case "modify":
						modify(change);
						break;
				}
			}
		}
	};
})(window);
`
