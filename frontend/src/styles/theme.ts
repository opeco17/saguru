import { grey, purple } from '@mui/material/colors';
import { createTheme } from '@mui/material/styles';

interface Color {
  main: string;
  light: string;
  dark: string;
  contrastText: string;
}

interface ColorOptions {
  main?: string;
  light?: string;
  dark?: string;
  contrastText?: string;
}

declare module '@mui/material/styles' {
  interface Palette {
    greyChip: Color;
    lightGreyChip: Color;
  }
  interface PaletteOptions {
    greyChip?: ColorOptions;
    lightGreyChip?: ColorOptions;
  }
}

const theme = createTheme({
  typography: {
    fontFamily: ['Roboto', 'Noto Sans JP'].join(','),
  },
  palette: {
    secondary: {
      main: purple['800'],
      light: purple['700'],
      dark: purple['900'],
      contrastText: '#fff',
    },
    info: {
      main: '#F85758',
      light: '#f57879',
      dark: '#F85758',
      contrastText: '#fff',
    },
    greyChip: {
      main: grey['800'],
      light: grey['700'],
      dark: grey['900'],
      contrastText: '#fff',
    },
    lightGreyChip: {
      main: grey['700'],
      light: grey['600'],
      dark: grey['800'],
      contrastText: '#fff',
    },
  },
});

export default theme;
