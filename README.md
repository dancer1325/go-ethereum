## Welcome to the go-ethereum website!

## Contributing

* issues
  * [go-ethereum repository / title's prefix `[website]: `](https://github.com/ethereum/go-ethereum/issues?q=is%3Aissue%20state%3Aopen%20%5Bwebsite%5D) 

## Stack

* == [Next.js](https://nextjs.org/) project / based on
  - [Node.js](https://nodejs.org/)
  - [React](https://reactjs.org/)
  - [Typescript](https://www.typescriptlang.org/)
  - [Chakra UI](https://chakra-ui.com/)
  - [Algolia](https://www.algolia.com/)
    - used for
      - Site indexing,
      - rapid intra-site search results,
      - search analytics
  - [Netlify](https://www.netlify.com/)
    - used for
      - DNS management
      - primary host -- for -- `master` build

## Repository structure

* [content](docs)
* [Website code](src)

## how to add a NEW documentation page?

* steps
  * | [docs](docs), 
    * add the NEW .md
  * | ["index.md"](docs/index.md)
    * if the NEW file is a NEW subsection
  * | ["documentation-links.yaml"](src/data/documentation-links.yaml)
    * define the left sidebar

## how to add notes | doc?

* == highlighted boxes
* steps

    ```markdown
    <Note>text to include in note</Note>
    ```

<img width="809" alt="Screen Shot 2023-01-04 at 18 22 06" src="https://user-images.githubusercontent.com/948922/210652463-1fc0370e-815c-427d-9eff-64199a300460.png">

## how to add images?

* steps
  * place | "public/images/docs"
  * refer | markdown

    ```markdown
    ![alt-text](/images/docs/image-title.png)
    ```

## how to build locally?

* 
  ```bash
  npm run dev
  # or
  yarn dev
  ```
* | browser,
  * http://localhost:3000

