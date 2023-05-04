import { useLocale } from '../../hooks/locale';
import GitHubIcon from '@mui/icons-material/GitHub';
import TranslateIcon from '@mui/icons-material/Translate';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Fade from '@mui/material/Fade';
import IconButton from '@mui/material/IconButton';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import Toolbar from '@mui/material/Toolbar';
import Tooltip from '@mui/material/Tooltip';
import { useRouter } from 'next/router';
import { useState, MouseEvent } from 'react';

type LanguageMenuItemProps = {
  onClick: any;
  text: string;
};

const LanguageMenuItem = ({ onClick, text }: LanguageMenuItemProps) => {
  return (
    <MenuItem onClick={onClick} sx={{ px: 2.5, py: 1.5 }}>
      {text}
    </MenuItem>
  );
};

const NavBar = () => {
  const router = useRouter();
  const { t } = useLocale();

  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const handleClick = (event: MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleCangeLocale = (locale: string) => {
    router.push('/', '/', { locale });
    setAnchorEl(null);
  };

  return (
    <AppBar position='sticky' color='inherit'>
      <Toolbar>
        <Box
          component='img'
          src='/logo.png'
          alt='logo'
          sx={{ height: { xs: '40px', sm: '43px' } }}
        />
        <div style={{ flexGrow: 1 }}></div>
        <Tooltip title={t.TRANSLATE_ICON_MESSAGE}>
          <IconButton sx={{ mr: { xs: 0.5, sm: 1 } }} onClick={handleClick}>
            <TranslateIcon sx={{ fontSize: { xs: '28px', sm: '31px' } }} />
          </IconButton>
        </Tooltip>
        <Menu
          anchorEl={anchorEl}
          open={Boolean(anchorEl)}
          onClose={() => {
            setAnchorEl(null);
          }}
          autoFocus={false}
          TransitionComponent={Fade}
        >
          <LanguageMenuItem onClick={() => handleCangeLocale('en')} text='English' />
          <LanguageMenuItem onClick={() => handleCangeLocale('ja')} text='Japanese' />
        </Menu>
        <Tooltip title={t.GITHUB_ICON_MESSAGE}>
          <a href='https://github.com/opeco17/saguru' target='_blank' rel='noreferrer'>
            <IconButton>
              <GitHubIcon sx={{ fontSize: { xs: '34px', sm: '37px' } }} />
            </IconButton>
          </a>
        </Tooltip>
      </Toolbar>
    </AppBar>
  );
};

export default NavBar;
