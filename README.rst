===================
Go Template Manager
===================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/gotemplateutil?status.png
   :target: https://godoc.org/github.com/siongui/gotemplateutil

.. image:: https://api.travis-ci.org/siongui/gotemplateutil.png?branch=master
   :target: https://travis-ci.org/siongui/gotemplateutil

.. image:: https://goreportcard.com/badge/github.com/siongui/gotemplateutil
   :target: https://goreportcard.com/report/github.com/siongui/gotemplateutil

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/gotemplateutil/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

Go_ tempalte manager (`text/template`_ and `html/template`_).

Development Environment:

  - `Ubuntu 16.10`_
  - `Go 1.7.5`_


Install
+++++++

.. code-block:: bash

  go get -u github.com/siongui/gotemplateutil


Idea (not implemented)
++++++++++++++++++++++

.. code-block:: go

  import "html/template"

  type TemplateManager struct {
  	StdTmpl	*template.Template
  }

  // include gettext by default
  tm := gotm.New(name, dir)

  // i18n
  tm.SetupMessagesDomain("locale/")

  tm.SetLocale("zh_TW")

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

.. _Go: https://golang.org/
.. _Ubuntu 16.10: http://releases.ubuntu.com/16.10/
.. _Go 1.7.5: https://golang.org/dl/
.. _git clone: https://www.google.com/search?q=git+clone
.. _text/template: https://golang.org/pkg/text/template/
.. _html/template: https://golang.org/pkg/html/template/
.. _UNLICENSE: http://unlicense.org/
