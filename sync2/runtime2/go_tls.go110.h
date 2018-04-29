// +build go1.10

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifdef GOARCH_arm
#define LR R14
#endif

#ifdef GOARCH_amd64
#define	get_tls(r)	MOVQ TLS, r
#define	g(r)	0(r)(TLS*1)
#endif

#ifdef GOARCH_amd64p32
#define	get_tls(r)	MOVL TLS, r
#define	g(r)	0(r)(TLS*1)
#endif

#ifdef GOARCH_386
#define	get_tls(r)	MOVL TLS, r
#define	g(r)	0(r)(TLS*1)
#endif

#define G_ID_OFFSET 152
#define G_M_OFFSET 48
#define M_P_OFFSET 208
#define P_ID_OFFSET 8