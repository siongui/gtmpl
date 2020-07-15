================================
Go Template With Gettext Support
================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/gtmpl?status.svg
   :target: https://godoc.org/github.com/siongui/gtmpl

.. image:: https://api.travis-ci.org/siongui/gotm.svg?branch=master
   :target: https://travis-ci.org/siongui/gotm

.. image:: https://goreportcard.com/badge/github.com/siongui/gtmpl
   :target: https://goreportcard.com/report/github.com/siongui/gtmpl

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/gtmpl/blob/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

Go_ `html/template`_ with gettext_ support.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.14.4`_


Install
+++++++

.. code-block:: bash

  go get -u github.com/siongui/gtmpl


Usage
+++++

The gettext_ feature is provided by `chai2010/gettext-go`_, which is pure Go
solution for gettext_.

.. code-block:: go

  import (
  	"github.com/siongui/gtmpl"
  )

  // TODO: add sample code here


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `Golang Template Inheritance (Python Jinja2 extends & include) <https://siongui.github.io/2017/02/05/go-template-inheritance-jinja2-extends-include/>`_

.. [2] `GitHub - chai2010/gettext-go: GNU gettext for Go <https://github.com/chai2010/gettext-go>`_
       (`godoc <https://godoc.org/github.com/chai2010/gettext-go/gettext>`__)

.. [3] `io writer to string golang - Google search <https://www.google.com/search?q=io+writer+to+string+golang>`_

.. [4] | `godoc code - Google search <https://www.google.com/search?q=godoc+code>`_
       | `godoctricks - GoDoc <https://godoc.org/github.com/fluhus/godoc-tricks>`_

.. [5] | render_ and grender_
       | `grender: stdlib templates with additional "extends" support : golang <https://www.reddit.com/r/golang/comments/61hcfg/grender_stdlib_templates_with_additional_extends/>`_
       | `GitHub - dannyvankooten/grender: Go package for easily rendering JSON/XML data and HTML templates <https://github.com/dannyvankooten/grender>`_
       | `GitHub - unrolled/render: Go package for easily rendering JSON, XML, binary data, and HTML templates responses. <https://github.com/unrolled/render>`_

.. [6] | `Go template rendering or Node.js template rendering? : golang <https://www.reddit.com/r/golang/comments/68i04t/go_template_rendering_or_nodejs_template_rendering/>`_
       | `GitHub - titpetric/egon: An ERB-style templating language for Go. <https://github.com/titpetric/egon>`_
       | `books/12fa-docker-golang/chapter7 at master · titpetric/books · GitHub <https://github.com/titpetric/books/tree/master/12fa-docker-golang/chapter7>`_

.. _Go: https://golang.org/
.. _grender: https://github.com/dannyvankooten/grender
.. _render: https://github.com/unrolled/render
.. _gettext: https://www.google.com/search?q=gettext
.. _chai2010/gettext-go: https://github.com/chai2010/gettext-go
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.14.4: https://golang.org/dl/
.. _git clone: https://www.google.com/search?q=git+clone
.. _text/template: https://golang.org/pkg/text/template/
.. _html/template: https://golang.org/pkg/html/template/
.. _UNLICENSE: https://unlicense.org/
