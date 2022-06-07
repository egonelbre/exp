# Code from:
# https://raw.githubusercontent.com/llimllib/pymag-trees/master/buchheim.py

from gen import Tree, gentree
from operator import lt, gt
from sys import stdout

class DrawTree(object):
    def __init__(self, tree, parent=None, depth=0, number=1):
        self.x = -1.
        self.y = depth
        self.tree = tree
        self.children = [DrawTree(c, self, depth+1, i+1)
                         for i, c
                         in enumerate(tree.children)]
        self.parent = parent
        self.thread = None
        self.mod = 0
        self.ancestor = self
        self.change = self.shift = 0
        self._lmost_sibling = None
        #this is the number of the node in its group of siblings 1..n
        self.number = number

    def left(self): 
        return self.thread or len(self.children) and self.children[0]

    def right(self):
        return self.thread or len(self.children) and self.children[-1]

    def lbrother(self):
        n = None
        if self.parent:
            for node in self.parent.children:
                if node == self: return n
                else:            n = node
        return n

    def get_lmost_sibling(self):
        if not self._lmost_sibling and self.parent and self != self.parent.children[0]:
            self._lmost_sibling = self.parent.children[0]
        return self._lmost_sibling
    lmost_sibling = property(get_lmost_sibling)

    def __str__(self): return "%s: x=%s mod=%s" % (self.tree, self.x, self.mod)
    def __repr__(self): return self.__str__()

def buchheim(tree):
    dt = firstwalk(DrawTree(tree))
    min = second_walk(dt)
    if min < 0:
        third_walk(dt, -min)
    return dt

def third_walk(tree, n):
    tree.x += n
    for c in tree.children:
        third_walk(c, n)

def firstwalk(n, distance=1.):
    if len(n.children) == 0:
        if n.lmost_sibling:
            n.x = n.lbrother().x + distance
        else:
            n.x = 0.
    else:
        default_ancestor = n.children[0]
        for c in n.children:
            firstwalk(c)
            default_ancestor = apportion(c, default_ancestor, distance)
        print( "finished n =", n.tree, "children")
        execute_shifts(n)

        midpoint = (n.children[0].x + n.children[-1].x) / 2

        ell = n.children[0]
        arr = n.children[-1]
        c = n.lbrother()
        if c:
            n.x = c.x + distance
            n.mod = n.x - midpoint
        else:
            n.x = midpoint
    return n

def apportion(n, default_ancestor, distance):
    c = n.lbrother()
    if c is None:
        return default_ancestor

    #in buchheim notation:
    #i == inner; o == outer; r == right; l == left; r = +; l = -
    vir = vor = n
    vil = c
    vol = n.lmost_sibling
    sir = sor = n.mod
    sil = vil.mod
    sol = vol.mod
    while vil.right() and vir.left():
        vil = vil.right()
        vir = vir.left()
        vol = vol.left()
        vor = vor.right()
        vor.ancestor = n
        shift = (vil.x + sil) - (vir.x + sir) + distance
        if shift > 0:
            move_subtree(ancestor(vil, n, default_ancestor), n, shift)
            sir = sir + shift
            sor = sor + shift
        sil += vil.mod
        sir += vir.mod
        sol += vol.mod
        sor += vor.mod
    if vil.right() and not vor.right():
        vor.thread = vil.right()
        vor.mod += sil - sor
    else:
        if vir.left() and not vol.left():
            vol.thread = vir.left()
            vol.mod += sir - sol
        default_ancestor = n
    return default_ancestor

def move_subtree(wl, wr, shift):
    subtrees = wr.number - wl.number
    print( wl.tree, "is conflicted with", wr.tree, 'moving', subtrees, 'shift', shift)
    #print( wl, wr, wr.number, wl.number, shift, subtrees, shift/subtrees)
    wr.change -= shift / subtrees
    wr.shift += shift
    wl.change += shift / subtrees
    wr.x += shift
    wr.mod += shift

def execute_shifts(n):
    shift = change = 0
    for c in n.children[::-1]:
        print( "shift:", c, shift, c.change)
        c.x += shift
        c.mod += shift
        change += c.change
        shift += c.shift + change

def ancestor(vil, n, default_ancestor):
    #the relevant text is at the bottom of page 7 of
    #"Improving Walker's Algorithm to Run in Linear Time" by Buchheim et al, (2002)
    #http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.16.8757&rep=rep1&type=pdf
    if vil.ancestor in n.parent.children:
        return vil.ancestor
    else:
        return default_ancestor

def second_walk(n, m=0, depth=0, min=None):
    n.x += m
    n.y = depth

    if min is None or n.x < min:
        min = n.x

    for c in n.children:
        min = second_walk(c, m + n.mod, depth+1, min)

    return min

r = 30
rh = r*1.5
rw = r*1.5
stroke(0)

def drawt(root, depth):
    global r
    oval(root.x * rw, depth * rh, r, r)
    print( root.x)
    for child in root.children:
        drawt(child, depth+1)

def drawconn(root, depth):
    for child in root.children:
        line(root.x * rw + (r/2), depth * rh + (r/2),
             child.x * rw + (r/2), (depth+1) * rh + (r/2))
        drawconn(child, depth+1)

size(1000, 500)
translate(2, 2)
stroke(0)
drawconn(dt, 0)
fill(1,1,1)
drawt(dt, 0)