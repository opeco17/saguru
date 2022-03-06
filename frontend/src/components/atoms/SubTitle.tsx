import { useLocale } from '../../hooks/locale';
import Typography from '@mui/material/Typography';
import { grey } from '@mui/material/colors';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';

const SubTitle = () => {
  const { t } = useLocale();
  const theme = useTheme();
  const isXSmall = useMediaQuery(theme.breakpoints.only('xs'));

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
      <span style={{ color: theme.palette.info.main }}>{t.SUB_TITLE_PREFIX}</span>
      <span style={{ color: grey['700'] }}>{t.SUB_TITLE}</span>
    </Typography>
  );
};

export default SubTitle;
