import Typography from '@mui/material/Typography';
import { useTheme } from '@mui/material/styles';
import { ReactNode } from 'react';

type InputTextProps = {
  children: ReactNode;
};

const InputText = ({ children }: InputTextProps) => {
  const theme = useTheme();

  return (
    <Typography
      variant='subtitle2'
      sx={{
        color: theme.palette.greyText.main,
      }}
    >
      {children}
    </Typography>
  );
};

export default InputText;
