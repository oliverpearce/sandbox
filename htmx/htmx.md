# htmx!

htmx aims to replace js frameworks with html.

[**HATEOA**](https://htmx.org/essays/hateoas/) (*hypermedia as the engine of application state*) adds attributes to html in order to handle complex ui requirements.
allows requests to be given to server from ANY element! 

this is done through a HTTP verb and the URL endpoint of the server:
```html
<button hx-get="https://google.com" hx-swap="outerHTML">
    click me!
</button>
```

- - -
## helpful links:
[documentation](https://htmx.org/docs/)