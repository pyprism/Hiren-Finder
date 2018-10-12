========
Overview
========

.. start-badges

.. list-table::
    :stub-columns: 1

    * - docs
      - |docs|
    * - tests
      - | |travis| |requires|
        | |coveralls|
    * - package
      - | |version| |wheel| |supported-versions| |supported-implementations|
        | |commits-since|



.. |travis| image:: https://travis-ci.org/pyprism/isBot.svg?branch=master
    :alt: Travis-CI Build Status
    :target: https://travis-ci.org/pyprism/isBot

.. |requires| image:: https://requires.io/github/pyprism/isBot/requirements.svg?branch=master
    :alt: Requirements Status
    :target: https://requires.io/github/pyprism/isBot/requirements/?branch=master

.. |coveralls| image:: https://coveralls.io/repos/pyprism/isBot/badge.svg?branch=master&service=github
    :alt: Coverage Status
    :target: https://coveralls.io/r/pyprism/isBot

.. |version| image:: https://img.shields.io/pypi/v/isBot.svg
    :alt: PyPI Package latest release
    :target: https://pypi.python.org/pypi/isBot

.. |commits-since| image:: https://img.shields.io/github/commits-since/pyprism/isBot/v1.0.0.svg
    :alt: Commits since latest release
    :target: https://github.com/pyprism/isBot/compare/v1.0.0...master

.. |wheel| image:: https://img.shields.io/pypi/wheel/isBot.svg
    :alt: PyPI Wheel
    :target: https://pypi.python.org/pypi/isBot

.. |supported-versions| image:: https://img.shields.io/pypi/pyversions/isBot.svg
    :alt: Supported versions
    :target: https://pypi.python.org/pypi/isBot

.. |supported-implementations| image:: https://img.shields.io/pypi/implementation/isBot.svg
    :alt: Supported implementations
    :target: https://pypi.python.org/pypi/isBot


.. end-badges

Detects crawler via user agent

* Free software: MIT license

Installation
============

::

    pip install isBot

Documentation
=============


To use the project:

.. code-block:: python

    import isBot
    isBot.longest()


Development
===========

To run the all tests run::

    tox

Note, to combine the coverage data from all the tox environments run:

.. list-table::
    :widths: 10 90
    :stub-columns: 1

    - - Windows
      - ::

            set PYTEST_ADDOPTS=--cov-append
            tox

    - - Other
      - ::

            PYTEST_ADDOPTS=--cov-append tox
