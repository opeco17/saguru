import { useLocale } from '../../hooks/locale';
import SimpleSelectWrapper from '../atoms/SimpleSelectWrapper';
import MenuItem from '@mui/material/MenuItem';
import Typography from '@mui/material/Typography';

type SimpleSelectFieldProps = {
  value: string;
  items: string[];
  onChange: any;
};

const SimpleSelectField = (props: SimpleSelectFieldProps) => {
  const { t } = useLocale();

  return (
    <>
      <SimpleSelectWrapper value={props.value} onChange={props.onChange} items={props.items}>
        {props.items.map((each) => {
          return (
            <MenuItem value={each} key={each}>
              <Typography>{t[each]}</Typography>
            </MenuItem>
          );
        })}
      </SimpleSelectWrapper>
    </>
  );
};

export default SimpleSelectField;
