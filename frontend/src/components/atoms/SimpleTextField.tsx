import TextField from '@mui/material/TextField';
import { useTheme } from '@mui/material/styles';

type SimpleTextFieldProps = {
  value: string;
  placeholder: string;
  onChange: (event: any) => void;
};

const SimpleTextField = ({ value, placeholder, onChange }: SimpleTextFieldProps) => {
  const theme = useTheme();

  return (
    <>
      <TextField
        size='small'
        placeholder={placeholder}
        value={value}
        onChange={onChange}
        sx={{ width: '100%' }}
        InputProps={{
          sx: { color: theme.palette.greyText.main },
        }}
      ></TextField>
    </>
  );
};

export default SimpleTextField;
