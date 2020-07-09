# Assistant Design Doc

This is how Assistant achieves its purpose.

## Event Loop

The main part of the software is the event loop which provides a two methods.

* `Emit`: Dispatches a new event.
* `On`: Registers a new handler for a specific event.

## Event Generators

Obviously we need a set of components in charge of producing new events.

## Event Handlers

We also need a set of components receive and process the generated events.

## Input Plugins

Input plugins help us to retrieve inforamtion from outside of the system and feed it into event generators or handlers.

## Output Plugins

Output plugins help us to export information to other systems from event generators or handlers.

## Preprocessors

Preprocessors are utility functions to manipulate data within event generators or handlers.
