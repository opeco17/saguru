import { Html, Head, Main, NextScript } from 'next/document';

const Document = () => {
  return (
    <Html>
      <Head>
        <link
          rel='stylesheet'
          href='https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap'
        />
        <link
          href='https://fonts.googleapis.com/css2?family=Noto+Sans+JP:wght@400;500&display=swap'
          rel='stylesheet'
        />
        <link rel='canonical' href='https://saguru.opeco17.com'></link>
        <link rel='shortcut icon' type='image/x-icon' href='favicon.ico' />
      </Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
};

export default Document;
