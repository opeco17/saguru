import { FormControl } from '@mui/material';
import { SelectChangeEvent } from '@mui/material';
import Select from '@mui/material/Select';
import { ReactNode } from 'react';

type SimpleSelectWrapperProps = {
  value: string;
  onChange: (event: SelectChangeEvent<string>) => void;
  children: ReactNode;
};

const SimpleSelectWrapper = ({ value, onChange, children }: SimpleSelectWrapperProps) => {
  return (
    <FormControl fullWidth>
      <Select
        size='small'
        value={value}
        onChange={onChange}
        sx={{ width: '100%' }}
        MenuProps={{
          sx: { maxHeight: 300 },
          anchorOrigin: { vertical: 'bottom', horizontal: 'left' },
          transformOrigin: { vertical: 'top', horizontal: 'left' },
        }}
      >
        {children}
      </Select>
    </FormControl>
  );
};

export default SimpleSelectWrapper;
