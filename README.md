# Glamour-slim

> [!NOTE]
>
> This is a permanent fork of [`charmbracelet/glamour`](https://github.com/charmbracelet/glamour) removing its 14MB dependency on [`yuin/goldmark-emoji`](https://github.com/yuin/goldmark-emoji).
>
> If your project has an indirect dependency on `charmbracelet/glamour`, you can decrease its size by adding the following to your `go.mod`:
>
> ```plaintext
> replace github.com/charmbracelet/glamour => github.com/justenstall/glamour-slim v0.0.0
> ```

Stylesheet-based markdown rendering for your CLI apps.

![Glamour dark style example](https://stuff.charm.sh/glamour/glamour-example.png)

`glamour` lets you render [markdown](https://en.wikipedia.org/wiki/Markdown)
documents & templates on [ANSI](https://en.wikipedia.org/wiki/ANSI_escape_code)
compatible terminals. You can create your own stylesheet or simply use one of
the stylish defaults.

## Usage

```go
import "github.com/justenstall/glamour-slim"

in := `# Hello World

This is a simple example of Markdown rendering with Glamour!
Check out the [other examples](https://github.com/charmbracelet/glamour/tree/master/examples) too.

Bye!
`

out, err := glamour.Render(in, "dark")
fmt.Print(out)
```

<!-- markdownlint-disable-next-line no-inline-html -->
<img src="https://github.com/charmbracelet/glamour/raw/master/examples/helloworld/helloworld.png" width="600" alt="Hello World example">

### Custom Renderer

```go
import "github.com/justenstall/glamour-slim"

r, _ := glamour.NewTermRenderer(
    // detect background color and pick either the default dark or light theme
    glamour.WithAutoStyle(),
    // wrap output at specific width (default is 80)
    glamour.WithWordWrap(40),
)

out, err := r.Render(in)
fmt.Print(out)
```

## Styles

You can find all available default styles in our [gallery](https://github.com/charmbracelet/glamour/tree/master/styles/gallery).
Want to create your own style? [Learn how!](https://github.com/charmbracelet/glamour/tree/master/styles)

There are a few options for using a custom style:

1. Call `glamour.Render(inputText, "desiredStyle")`
1. Set the `GLAMOUR_STYLE` environment variable to your desired default style or a file location for a style and call `glamour.RenderWithEnvironmentConfig(inputText)`
1. Set the `GLAMOUR_STYLE` environment variable and pass `glamour.WithEnvironmentConfig()` to your custom renderer

## Glamourous Projects

Check out these projects, which use `glamour`:

- [Glow](https://github.com/charmbracelet/glow), a markdown renderer for
the command-line.
- [GitHub CLI](https://github.com/cli/cli), GitHub’s official command line tool.
- [GitLab CLI](https://gitlab.com/gitlab-org/cli), GitLab's official command line tool.
- [Gitea CLI](https://gitea.com/gitea/tea), Gitea's official command line tool.
- [Meteor](https://github.com/odpf/meteor), an easy-to-use, plugin-driven metadata collection framework.

## Feedback

We’d love to hear your thoughts on this project. Feel free to drop us a note!

- [Twitter](https://twitter.com/charmcli)
- [The Fediverse](https://mastodon.social/@charmcli)
- [Discord](https://charm.sh/chat)

## License

[MIT](https://github.com/charmbracelet/glamour/raw/master/LICENSE)

***

Part of [Charm](https://charm.sh).

<!-- markdownlint-disable-next-line no-inline-html -->
<a href="https://charm.sh/"><img alt="The Charm logo" src="https://stuff.charm.sh/charm-badge.jpg" width="400"></a>

Charm热爱开源 • Charm loves open source
