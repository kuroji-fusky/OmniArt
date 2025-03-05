# OmniArt

An all-in-one solution to search artworks from known art platforms!

## Supported platforms

> [!NOTE]
>
> All requests are cached to circument rate-limiting and stay persistent for an hour.
> The hosted version will be limited to save on database hosting costs.

| Platform      | Method    | API Access                 |
| ------------- | --------- | -------------------------- |
| ArtStation    | API       | Yes, via internal requests |
| Behance       | API       | Yes                        |
| DeviantArt    | Web crawl | Limited                    |
| e621/e926     | Web crawl | No                         |
| FurAffinity   | Web crawl | No                         |
| Furry Network | API       | Yes, via internal requests |
| Giphy         | API       | Yes, requires auth         |
| InkBunny      | Web crawl | No                         |
| Itaku         | Web crawl | Unknown                    |
| Newgrounds    | Web crawl | No                         |
| Pinterest     | API       | Yes, requires auth         |
| Pixiv         | API       | Limited                    |
| Tenor         | API       | Yes, requires auth         |
| Toyhouse      | Web crawl | No                         |
| Weasyl        | Web crawl | No                         |

Unfortunately, platforms such as Twitter/X and Instagram are excluded because they're notoriously difficult to scrape and have limited API access. With Twitter having a paywalled access to their API isn't feasble for those who prefer to deploy through Docker.

They can be ambiguous for providing a simple search query, and having to create additional complex filters by handle can already put a lot of technical burden on this project.

However, if you wish to be excluded from the list, feel free to reach me out at hello@kurojifusky[.]com.

## Contributing

### Setup

A latest version of Node.js or the latest LTS version is installed on your system.

1. Install its dependencies using `pnpm i`
1. Generate an API key with `pnpm run generate-key`, or alternatively, `node scripts/generate-key.js`, it should display out:

   ```console
   Token generated! This token has already populated from an .env file.

   om_<random 96 characters>

   ```

   It will add this key automatically to your `.env` file

## License

[GPL-2.0](/LICENSE)
