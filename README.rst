===================
Go Template Manager
===================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/gotm?status.png
   :target: https://godoc.org/github.com/siongui/gotm

.. image:: https://api.travis-ci.org/siongui/gotm.png?branch=master
   :target: https://travis-ci.org/siongui/gotm

.. image:: https://goreportcard.com/badge/github.com/siongui/gotm
   :target: https://goreportcard.com/report/github.com/siongui/gotm

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/gotm/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

Go_ tempalte manager (`text/template`_ and `html/template`_). Try to be similar
to grender_ or render_, but with gettext_ support in templates.

Development Environment:

  - `Ubuntu 16.10`_
  - `Go 1.8.1`_


Install
+++++++

.. code-block:: bash

  go get -u github.com/siongui/gotm


Idea (not finish to implement)
++++++++++++++++++++++++++++++

.. code-block:: go

  import "github.com/siongui/gotm"

  // include gettext by default
  tm := gotm.New(name, dir)

  // i18n
  gotm.SetupMessagesDomain("locale/")

  gotm.SetLocale("zh_TW")

  // output
  tm.Render(io.Writer, tmplname, yamlfile)


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

.. [5] | `grender: stdlib templates with additional "extends" support : golang <https://www.reddit.com/r/golang/comments/61hcfg/grender_stdlib_templates_with_additional_extends/>`_
       | `GitHub - dannyvankooten/grender: Go package for easily rendering JSON/XML data and HTML templates <https://github.com/dannyvankooten/grender>`_
       | `GitHub - unrolled/render: Go package for easily rendering JSON, XML, binary data, and HTML templates responses. <https://github.com/unrolled/render>`_

.. [6] | `Go template rendering or Node.js template rendering? : golang <https://www.reddit.com/r/golang/comments/68i04t/go_template_rendering_or_nodejs_template_rendering/>`_
       | `GitHub - titpetric/egon: An ERB-style templating language for Go. <https://github.com/titpetric/egon>`_
       | `books/12fa-docker-golang/chapter7 at master · titpetric/books · GitHub <https://github.com/titpetric/books/tree/master/12fa-docker-golang/chapter7>`_

.. _Go: https://golang.org/
.. _grender: https://github.com/dannyvankooten/grender
.. _render: https://github.com/unrolled/render
.. _gettext: https://github.com/chai2010/gettext-go
.. _Ubuntu 16.10: http://releases.ubuntu.com/16.10/
.. _Go 1.8.1: https://golang.org/dl/
.. _git clone: https://www.google.com/search?q=git+clone
.. _text/template: https://golang.org/pkg/text/template/
.. _html/template: https://golang.org/pkg/html/template/
.. _UNLICENSE: http://unlicense.org/
