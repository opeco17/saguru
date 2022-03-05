import { FormControl } from '@mui/material';
import Select from '@mui/material/Select';
import { ReactNode } from 'react';

type SimpleSelectWrapperProps = {
  value: string;
  onChange: any;
  items: (string | boolean)[];
  children: ReactNode;
};

const SimpleSelectWrapper = (props: SimpleSelectWrapperProps) => {
  return (
    <>
      <FormControl fullWidth>
        <Select
          size='small'
          value={props.value}
          onChange={props.onChange}
          sx={{ width: '100%' }}
          MenuProps={{
            sx: { maxHeight: 300 },
            anchorOrigin: { vertical: 'bottom', horizontal: 'left' },
            transformOrigin: { vertical: 'top', horizontal: 'left' },
          }}
        >
          {props.children}
        </Select>
      </FormControl>
    </>
  );
};

export default SimpleSelectWrapper;
