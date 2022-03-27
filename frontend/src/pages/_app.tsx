import GoogleAnalytics from '../components/atoms/GoogleAnalytics';
import { GA_TRACKING_ID } from '../lib/gtag';
import '../styles/globals.css';
import theme from '../styles/theme';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import type { AppProps } from 'next/app';

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      {GA_TRACKING_ID !== '' && <GoogleAnalytics />}
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Component {...pageProps} />
      </ThemeProvider>
    </>
  );
};

export default App;
