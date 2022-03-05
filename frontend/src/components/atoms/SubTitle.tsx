import { useLocale } from '../../hooks/locale';
import Typography from '@mui/material/Typography';
import { grey } from '@mui/material/colors';
import { Theme, useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';

const SubTitle = () => {
  const { t } = useLocale();
  const theme = useTheme();
  const isXSmall = useMediaQuery((theme: Theme) => theme.breakpoints.only('xs'));
  const greyColor = grey['700'];
  const infoColor = theme.palette.info.main;

  return (
    <Typography
      variant={isXSmall ? 'h6' : 'h5'}
      component='h2'
      sx={{
        fontWeight: { xs: 'medium' },
        lineHeight: '1.5',
        textAlign: 'left',
        display: 'inline-block',
      }}
    >
      <span style={{ color: infoColor }}>{t.SUB_TITLE_PREFIX}</span>
      <span style={{ color: greyColor }}>{t.SUB_TITLE}</span>
    </Typography>
  );
};

export default SubTitle;
