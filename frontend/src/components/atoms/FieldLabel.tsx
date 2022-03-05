import Typography from '@mui/material/Typography';
import { grey } from '@mui/material/colors';
import { ReactNode } from 'react';

type FieldLabelProps = {
  children: ReactNode;
};

const FieldLabel = (props: FieldLabelProps) => {
  const greyColor = grey['700'];

  return (
    <Typography
      sx={{
        textAlign: 'left',
        color: greyColor,
        fontSize: 15,
        mb: 0.5,
      }}
    >
      {props.children}
    </Typography>
  );
};

export default FieldLabel;
