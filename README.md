# krakend-martian-correlation-header
This plugin is useful if you need generate some kind of correlation id

## Using it as a lib
Import the required package into your project and start using them as any other martian component

```
import(
	_ "github.com/JaSei/krakend-martian-correlation-header/martian"
)
```

## Using the modules as KrakenD plugins
Compile the desired package with the plugin flag

```
$ go build -buildmode=plugin -o krakend-martian_correlation_header.so ./krakend-plugin/correlation_header
```
And place the plugins into your plugin folder, so the KrakenD can load them in runtime.

## Sample DSL

```
{
    "header.Correlation": {
        "header_name" : "X-Correlation-ID"
    }
}
```
