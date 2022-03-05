import { useLocale } from '../../hooks/locale';
import SimpleSelectWrapper from '../atoms/SimpleSelectWrapper';
import MenuItem from '@mui/material/MenuItem';
import Typography from '@mui/material/Typography';

type LicenseFieldProps = {
  value: string;
  onChange: any;
  items: string[];
};

const LicenseField = (props: LicenseFieldProps) => {
  const { t } = useLocale();
  return (
    <>
      <SimpleSelectWrapper value={props.value} onChange={props.onChange} items={props.items}>
        {props.items.map((each) => {
          return (
            <MenuItem value={each} key={each}>
              <Typography>{each === 'ALL' ? t.ALL : each}</Typography>
            </MenuItem>
          );
        })}
      </SimpleSelectWrapper>
    </>
  );
};

export default LicenseField;
