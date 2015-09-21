package plt

import (
    "github.com/z0mbie42/py3"
)

type pltCtx struct {
    ft_bar      *py3.PyObject
    ft_plot     *py3.PyObject
    ft_scatter  *py3.PyObject
    ft_show     *py3.PyObject
    module      *py3.PyObject
}

var ctx pltCtx

func init() {
	py3.Initialize()
    ctx.module = py3.ImportModule("matplotlib.pyplot")
    ctx.ft_bar = ctx.module.GetAttrString("bar")
    ctx.ft_plot = ctx.module.GetAttrString("plot")
    ctx.ft_scatter = ctx.module.GetAttrString("scatter")
    ctx.ft_show = ctx.module.GetAttrString("show")
}

func Bar(x, y []float64) () {
    xPy := py3.NewList(len(x))
    yPy := py3.NewList(len(y))

    for i, xx := range x {
        xPy.SetItem(i, py3.PyFloat_FromDouble(xx))
        yPy.SetItem(i, py3.PyFloat_FromDouble(y[i]))
    }

    args := py3.NewTuple(2)
    args.SetItem(0, xPy)
    args.SetItem(1, yPy)

    ctx.ft_bar.CallObject(args)

    xPy.PyObj().DecRef()
    yPy.PyObj().DecRef()
    args.PyObj().DecRef()
}

func Plot(x, y []float64) () {

    xPy := py3.NewList(len(x))
    yPy := py3.NewList(len(y))

    for i, xx := range x {
        xPy.SetItem(i, py3.PyFloat_FromDouble(xx))
        yPy.SetItem(i, py3.PyFloat_FromDouble(y[i]))
    }

    args := py3.NewTuple(2)
    args.SetItem(0, xPy)
    args.SetItem(1, yPy)

    ctx.ft_plot.CallObject(args)

    xPy.PyObj().DecRef()
    yPy.PyObj().DecRef()
    args.PyObj().DecRef()
}

func Scatter(x, y []float64) () {
    xPy := py3.NewList(len(x))
    yPy := py3.NewList(len(y))

    for i, xx := range x {
        xPy.SetItem(i, py3.PyFloat_FromDouble(xx))
        yPy.SetItem(i, py3.PyFloat_FromDouble(y[i]))
    }

    args := py3.NewTuple(2)
    args.SetItem(0, xPy)
    args.SetItem(1, yPy)

    ctx.ft_scatter.CallObject(args)

    xPy.PyObj().DecRef()
    yPy.PyObj().DecRef()
}

func Show() () {
	empty_tuple := py3.NewTuple(0)
    ctx.ft_show.CallObject(empty_tuple)
    empty_tuple.PyObj().DecRef()
}
