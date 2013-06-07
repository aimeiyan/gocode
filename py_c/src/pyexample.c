/* pyexample.c */
#include "Python.h"
#include "example.h"
static char py_gcd_doc[] = "Computes the GCD of two integers";
static PyObject * py_gcd(PyObject *self, PyObject *args) {
    int x,y,r;
    if (!PyArg_ParseTuple(args,"ii:gcd",&x,&y)) {
        return NULL;
    }
    r = gcd(x,y);
    return Py_BuildValue("i",r);
}
static char py_replace_doc[] = "Replaces all characters in a string";
static PyObject * py_replace(PyObject *self, PyObject *args, PyObject *kwargs) {
    static char *argnames[] = {"s","och","nch",NULL};
    char *s,*sdup;
    char och, nch;
    int nrep;
    PyObject *result;
    if (!PyArg_ParseTupleAndKeywords(args,kwargs, "scc:replace",
                                     argnames, &s, &och, &nch)) {
        return NULL;
    }
    sdup = (char *) malloc(strlen(s)+1);
    strcpy(sdup,s);
    nrep = replace(sdup,och,nch);
    result = Py_BuildValue("(is)",nrep,sdup);
    free(sdup);
    return result;
}
static char py_distance_doc[] = "Computes the distance between two points";
static PyObject * py_distance(PyObject *self, PyObject *args) {
    PyErr_SetString(PyExc_NotImplementedError,"distance() not implemented.");
    return NULL;
}
static PyMethodDef _examplemethods[] = {
    {"gcd", py_gcd, METH_VARARGS, py_gcd_doc},
    {"replace", py_replace, METH_VARARGS | METH_KEYWORDS, py_replace_doc},
    {"distance",py_distance,METH_VARARGS, py_distance_doc},
    {NULL, NULL, 0, NULL}
};
#if PY_MAJOR_VERSION < 3
/* Python 2 module initialization */
void init_example(void) {
    PyObject *mod;
    mod = Py_InitModule("_example", _examplemethods);
    PyModule_AddIntMacro(mod,MAGIC);
}
#else
/* Python 3 module initialization */
static struct PyModuleDef _examplemodule = {
    PyModuleDef_HEAD_INIT,
    "_example",
    /* name of module */
    NULL,
    /* module documentation, may be NULL */
    -1,
    _examplemethods
};
PyMODINIT_FUNC
PyInit_ _example(void) {
    PyObject *mod;
    mod = PyModule_Create(&_examplemodule);
    PyModule_AddIntMacro(mod, MAGIC);
    return mod;
}
#endif
