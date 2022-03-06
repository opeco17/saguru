import { useLocale } from '../../hooks/locale';
import SimpleSelectWrapper from '../atoms/SimpleSelectWrapper';
import { SelectChangeEvent } from '@mui/material';
import MenuItem from '@mui/material/MenuItem';
import Typography from '@mui/material/Typography';

type LicenseFieldProps = {
  value: string;
  items: string[];
  onChange: (event: SelectChangeEvent) => void;
};

const LicenseField = ({ value, onChange, items }: LicenseFieldProps) => {
  const { t } = useLocale();
  return (
    <SimpleSelectWrapper value={value} onChange={onChange}>
      {items.map((each) => {
        return (
          <MenuItem value={each} key={each}>
            <Typography>{each === 'ALL' ? t.ALL : each}</Typography>
          </MenuItem>
        );
      })}
    </SimpleSelectWrapper>
  );
};

export default LicenseField;
