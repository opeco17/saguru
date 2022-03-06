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
  values: string[];
  onChange: any;
  onChipDelete: any;
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
          value={props.values}
          onChange={props.onChange}
          MenuProps={{
            sx: { maxHeight: 330 },
            anchorOrigin: { vertical: 'bottom', horizontal: 'left' },
            transformOrigin: { vertical: 'top', horizontal: 'left' },
          }}
          renderValue={(values: readonly string[]) => (
            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
              {values.map((value, index) => (
                <Chip
                  variant='outlined'
                  size='small'
                  // @ts-ignore to use custom color
                  color='greyChip'
                  label={value === 'ALL' ? t.ALL : value}
                  key={index}
                  sx={{ px: 0.5 }}
                  onDelete={() => {
                    props.onChipDelete(props.values.filter((each) => each !== value));
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
                checked={props.values.indexOf(option) > -1}
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
