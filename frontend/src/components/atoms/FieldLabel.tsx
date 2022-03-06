import Typography from '@mui/material/Typography';
import { grey } from '@mui/material/colors';
import { ReactNode } from 'react';

type FieldLabelProps = {
  children: ReactNode;
};

const FieldLabel = ({ children }: FieldLabelProps) => {
  return (
    <Typography
      sx={{
        textAlign: 'left',
        color: grey['700'],
        fontSize: 15,
        mb: 0.5,
      }}
    >
      {children}
    </Typography>
  );
};

export default FieldLabel;
