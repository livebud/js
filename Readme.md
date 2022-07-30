# JS

The JS package

> Note: This package is still a work in progress. Please see the [issues](https://github.com/livebud/js/issues) for what needs to be done.

## Currently supported Javascript VMs

- [V8](https://github.com/rogchap/v8go)
- [Goja](https://github.com/dop251/goja)

## Goals

- **Swappable JS VMs**: The available VMs each have pros and cons. You should be able to swap VMs out based on your needs.
- **Consistent Runtime**: For each of these VMs, there should be consistent and well-tested globals (e.g. `console`, `setTimeout`, `URL`) that match the web's behavior as much as possible.

## Non-Goals

- **Secure sandbox for user-submitted Javascript**: To provide better performance, the context is shared across evaluations. This package is not built to evaluate user-submitted Javascript.
- **Support non-standard runtime APIs**: There's no plans to add APIs that are specific to certain runtime environments such as Cloudflare Workers, Deno, etc. There's no science to this, but more of a heuristic. The API should be available in the web and in most server-side environments like Node.js.

## License

MIT
