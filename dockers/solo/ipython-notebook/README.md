zaha/ipython-notebook
---------------------

A standalone IPython Notebook docker with mountable notebook volume.

Get it: `docker pull zaha/ipython-notebook`
Use it: `docker run -it --rm -p 8080:8080 -v $(pwd):/ipython/notebooks zaha/ipython-notebook`

Built from docker `zaha/python-ml` (see folder bases/python-ml)
