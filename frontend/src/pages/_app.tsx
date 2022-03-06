import '../styles/globals.css';
import theme from '../styles/theme';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import type { AppProps } from 'next/app';
import GoogleAnalytics from '../components/atoms/GoogleAnalytics';

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <GoogleAnalytics />
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Component {...pageProps} />
      </ThemeProvider>
    </>
  );
};

export default App;
