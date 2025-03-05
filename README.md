# OmniArt

An all-in-one solution to search artworks from known art platforms!

## About

__OmniArt__ is an open-source art aggregator website that allows you to search artworks from known art platforms, DeviantArt, Pixiv, Behance, in bulk!

- Filter out by searching from one or more [supported platforms](#supported-platforms)
- Ability to include explicit (or NSFW) content
- Includes flexible API
- CLI support soon!
- No ads and AI "algorithm" BS

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

Unfortunately, platforms such as Twitter/X and Instagram will never be included simply because they're notoriously difficult to scrape and have limited API access. With Twitter, in particular, having a paywalled access to their API, isn't feasble for those who prefer to deploy through Docker that already put strains the technical burden on this project.

## Contributing

### Setup

A latest version of Node.js or the latest LTS version is installed on your system.

1. Install its dependencies using `pnpm i`
1. Generate an API key with `pnpm run generate-key`, or alternatively, `node scripts/generate-key.js`, it should display out:

   ```console
   Token generated! This token has already populated from an `.env` file.

   om_<random 96 characters>

   ```

   It will add this key automatically to your `.env` file

## License

[GPL-2.0](/LICENSE)

## Disclaimer

OmniArt does not own any of the content displayed on its source code or from the website. All images and artworks remain the property of their respective copyright holders. OmniArt merely aggregates publicly available data for user's convenience and is intended solely for personal, non-commercial use.
