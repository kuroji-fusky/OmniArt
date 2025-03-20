# OmniArt

> [!NOTE]
>
> This project is still currently in active development, the final product still isn't available
> for the time being.

An all-in-one solution to search artworks from known art platforms!

## About

**OmniArt** is an open-source art aggregator website that allows you to search artworks from known art platforms, DeviantArt, Pixiv, Behance, etc. in bulk!

- Filter out by searching from one or more [supported platforms](#supported-platforms)
- Ability to include explicit (or NSFW) content
- Includes flexible API
- CLI support soon!
- Best of all: there are no invasive ads and AI "algorithm" BS

## Disclaimer

OmniArt does not own any of the content displayed on its source code or from the website. All images and artworks remain the property of their respective copyright holders. OmniArt merely aggregates publicly available data for user's convenience and is intended solely for personal, non-commercial use.

## Supported platforms

> [!NOTE]
>
> All requests are cached to circument rate-limiting and stay persistent for an hour.
> The hosted version will be limited to save on database hosting costs.

| Platform                                          | Method    | API Access                 |
| ------------------------------------------------- | --------- | -------------------------- |
| [ArtStation](https://artstation.com)              | API       | Yes, via internal requests |
| [Behance](https://behance.com)                    | API       | Yes                        |
| [DeviantArt](https://deviantart.com)              | Web crawl | Limited                    |
| [Dribbble](https://dribbble.com)                  | API       | Limited                    |
| [e621](https://e621.net)/[e926](https://e926.net) | Web crawl | No                         |
| [FurAffinity](https://furaffinity.net)            | Web crawl | No                         |
| [Furry Network](https://furrynetwork.com)         | API       | Yes, via internal requests |
| [Giphy](https://giphy.com)                        | API       | Yes, requires auth         |
| [InkBunny](https://inkbunny.com)                  | Web crawl | No                         |
| [Itaku](https://itaku.ee)                         | Web crawl | Unknown                    |
| [Newgrounds](https://newgrounds.com)              | Web crawl | No                         |
| [Pinterest](https://pinterest.com)                | API       | Yes, requires auth         |
| [Pixiv](https://pixiv.com)                        | API       | Limited                    |
| [Tenor](https://tenor.com)                        | API       | Yes, requires auth         |
| [Toyhouse](https://toyhou.se)                     | Web crawl | No                         |
| [Weasyl](https://weasyl.com)                      | Web crawl | No                         |

If you think there's something you'd like to add, feel free to create an issue! However, if you wish to be excluded from the list, feel free to reach me out at hello@kurojifusky[.]com.

### Limitations for Twitter/X and Instagram

Unfortunately, platforms such as Twitter/X and Instagram will _never_ be included; for Instagram, they're notoriously difficult to scrape and have limited API access. For Twitter in particular - with having a paywalled access to their APIs - isn't feasble for those who prefer to deploy through a Docker container strains the technical burden on this project.

## Contributing

After cloning/forking the project, you must have the proper prerequisites:

- Go (v1.22.0 or higher)
- Node.js (v20 or higher)
- PNPM package manager
- `make` (Optional)

### Setup

#### Installing manually

- Installing dependencies

  - In the `client` directory, run `pnpm install`, or its shorthand, `pnpm i`
  - In the `server` directory, run `go get download`

- Generating an API token

  - If you're already in the `server` directory, run the following command to auto-generate an `.env` file
    already populated with an API token for you

    ```sh
    go run cmd/generate-key
    ```

#### Installing with `make`

```bash
make
```

#### Via Docker

TBA

## License

[GPL-2.0](/LICENSE)
