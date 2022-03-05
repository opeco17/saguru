import { useLocale } from '../../hooks/locale';
import CheckBoxIcon from '@mui/icons-material/CheckBox';
import CheckBoxOutlineBlankIcon from '@mui/icons-material/CheckBoxOutlineBlank';
import { FormControl } from '@mui/material';
import Box from '@mui/material/Box';
import Checkbox from '@mui/material/Checkbox';
import Chip from '@mui/material/Chip';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';

type MultipleChipsSelectFieldProps = {
  options: string[];
  value: string[];
  onChange: any;
};

const icon = <CheckBoxOutlineBlankIcon fontSize='small' />;
const checkedIcon = <CheckBoxIcon fontSize='small' />;

const MultipleChipsSelectField = (props: MultipleChipsSelectFieldProps) => {
  const { t } = useLocale();
  return (
    <>
      <FormControl fullWidth>
        <Select
          size='medium'
          multiple
          sx={{ width: '100%' }}
          value={props.value}
          onChange={props.onChange}
          MenuProps={{
            sx: { maxHeight: 330 },
            anchorOrigin: { vertical: 'bottom', horizontal: 'left' },
            transformOrigin: { vertical: 'top', horizontal: 'left' },
          }}
          renderValue={(value: readonly string[]) => (
            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
              {value.map((each, index) => (
                <Chip
                  variant='outlined'
                  size='small'
                  // @ts-ignore to use custom color
                  color='greyChip'
                  label={each === 'ALL' ? t.ALL : each}
                  key={index}
                  sx={{ px: 0.5 }}
                  onDelete={() => {
                    alert(each);
                  }}
                  onMouseDown={(event: any) => {
                    event.stopPropagation();
                  }}
                />
              ))}
            </Box>
          )}
        >
          {props.options.map((option) => (
            <MenuItem key={option} value={option}>
              <Checkbox
                icon={icon}
                checkedIcon={checkedIcon}
                checked={props.value.indexOf(option) > -1}
              />
              {option}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </>
  );
};

export default MultipleChipsSelectField;
