# Extending Go Applications

The overall inspiration for the question is to be able to publish a
docker image with some base code, and allow additional functionality
by extending the Dockerfile.

In the `exec/` example, this would be
by compiling the external program `exec/plugins/plugin1/main.go` and adding it
to the config (not written) for `exec/main.g`.

```docker
FROM <base-app>

# build local plugin
RUN go build -o ./plugin1.x
# add local plugin to server
RUN echo './plugin1.x' >> some-config.yaml

# base-server.x (aka exec/main.go in example) runs with the plugin1 loaded.
CMD ["/base-server.x", "...args"]
```

At present, the plugins I'm imagining are data soruces that run once every 6 hours or so, rather
than updates to data served to clients during normal server operation.

Requirements I'm hoping to hit:

1. Allow users to easily add functionality
   - Easy will be entirely subjective, but should not require making changes to base server code
   - Added functionality will be in specific documented scope (getting events for a venue as in example)
2. Avoid needing to recompile base server

After writing out the following approaches, [The Exec](#the-exec-solution) looks like the quickest and most
flexible approach. Not the best design in terms of architecture, but it also doesn't feel like an over-engineered
solution.

## The `Exec` Solution

The example solution in `exec/` executes a new program and reads the JSON output
back into the main server.

This seems like the easiest to implement and would allow for people to add plugins
in different lagnuages, since the plugin just needs to writeback the correct JSON.

Spawning a new process seems like overkill, but it is the most flexible/easy solution
and the loads should be small enough that the performance issues shouldn't be a huge problem.

## Go's `plugin` Solution

The [plugin](https://pkg.go.dev/plugin) package allows for loading compiled go code
into a project. It sounds like there are a lot of caveats, but if the plugin is getting
built on the same image as the main server, maybe there is less of a problem.

## Web Assembly

This came up a lot in forums and might be a decent solution.

## Separate Service w/ API

Require the downstream developer to implement their own service with a defined API the
main server connects to.

The requires a lot more work on downstream developers, and is a lot more resource intensive.

It would also require an API spec for both the base server and the separate one.
