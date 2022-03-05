import '../styles/globals.css';
import theme from '../styles/theme';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import type { AppProps } from 'next/app';

const MyApp = ({ Component, pageProps }: AppProps) => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Component {...pageProps} />
    </ThemeProvider>
  );
};

export default MyApp;
