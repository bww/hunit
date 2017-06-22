# HUnit

HUnit is a command-line tool that runs tests against HTTP APIs.

Since tests and documentation are maintained in parallel, HUnit combines these two highly-related tasks into one. Descriptions can be added to your tests and documentation of your endpoints can be automatically generated.

## Running Tests

Tests can be run by pointing `hunit` to a test suite document.

```
$ hunit test.yml
  ====> test.yml
  ----> GET https://raw.githubusercontent.com/bww/hunit/master/example/test.txt
        #1 Unexpected status code:
             expected: (int) 404
               actual: (int) 200

        #2 Entities do not match:
             --- Expected
             +++ Actual
             @@ -1,2 +1,2 @@
             -Heres a simple
             +Here's a simple
              response from the
```

## Generating Documentation

When needed, documentation for the endpoints in your tests can be generated by passing the `-gendoc` flag. In this case tests are run and their input and output are automatically incorporated into documentation.

Refer to [`example/docs`](https://github.com/bww/hunit/blob/master/example/docs) for an example of the docs generated for the example tests.

```
$ hunit -gendoc test.yml
====> test.yml
# ...
```

## Describing Tests

Tests are described by a simple YAML-based document format. Tests and documentation are managed together.

The test suite document format is straightforward and highly readible. Advanced features are also available if you need them, like:

* Comparing entities semantically to ignore insignificant differences like whitespace,
* Referencing the output of previously-run tests from the currently-executing one,
* Some useful functions which can be invoked in string interpolation.

Refer to the full [`test.yml`](https://github.com/bww/hunit/blob/master/example/test.yml) file for a more complete illustration of test cases.

```yaml
- 
    doc: |
      Fetch a document from our [Github repo](github.com/bww/hunit).
    
    request:
      method: GET
      url: https://raw.githubusercontent.com/bww/hunit/master/example/test.txt
      headers:
        Origin: localhost
    
    response:
      status: 200
      headers:
        Content-Type: text/plain; charset=utf-8
      entity: |
        Heres a simple
        response from the
        server.      
```