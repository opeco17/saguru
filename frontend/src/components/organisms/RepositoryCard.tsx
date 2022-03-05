import { Repository } from '../../types/repository';
import RepositoryCardChip from '../atoms/RepositoryChip';
import IssueCard from './IssueCard';
import { mdiSourceFork } from '@mdi/js';
import Icon from '@mdi/react';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import StarBorderIcon from '@mui/icons-material/StarBorder';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Collapse from '@mui/material/Collapse';
import Divider from '@mui/material/Divider';
import Typography from '@mui/material/Typography';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useState } from 'react';

type RepositoryCardProps = {
  repository: Repository;
};

const RepositoryCard = (props: RepositoryCardProps) => {
  const theme = useTheme();
  const [isDetailOpen, setIsDetailOpen] = useState(false);
  const detailOpenIcon = isDetailOpen ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />;
  const formatKilo = (value: number): string =>
    value > 999 ? (value / 1000).toFixed(1).toString() + 'k' : value.toString();

  return (
    <Card variant='outlined' sx={{ mb: 2 }}>
      <CardContent sx={{ py: 1.5, '&:last-child': { paddingBottom: 1.5 } }}>
        <Typography
          variant='h4'
          component='div'
          sx={{ fontSize: 20, fontWeight: 'medium', mb: 1.5 }}
        >
          <Box
            component='a'
            href={props.repository.url}
            target='_blank'
            rel='noreferrer'
            sx={{ '&:hover': { color: theme.palette.info.main } }}
          >
            {props.repository.name}
          </Box>
        </Typography>
        <Typography
          variant='h6'
          color='text.secondary'
          sx={{ fontSize: 14, fontWeight: 'regular', mb: 2 }}
        >
          {props.repository.description}
        </Typography>
        <RepositoryCardChip label={`# ${props.repository.language}`} />
        <RepositoryCardChip
          label={formatKilo(props.repository.starCount)}
          icon={<StarBorderIcon sx={{ width: '0.8em' }} />}
        />
        {useMediaQuery(theme.breakpoints.up('md')) ? (
          <RepositoryCardChip
            label={formatKilo(props.repository.forkCount)}
            icon={<Icon path={mdiSourceFork} size={0.6} />}
          />
        ) : null}

        <Button
          variant='outlined'
          size='small'
          // @ts-ignore to use custom color
          color='secondary'
          sx={{ pl: 2, pr: 1, py: 0.1, float: 'right' }}
          onClick={() => setIsDetailOpen((prev) => !prev)}
        >
          {props.repository.issues.length} issues {detailOpenIcon}
        </Button>
        <Box sx={{ mb: 0.5 }}></Box>
        <Collapse orientation='vertical' in={isDetailOpen}>
          <Box sx={{ mt: 3 }}></Box>
          {props.repository.issues.map((issue) => {
            return (
              <Box key={issue.id}>
                <Divider />
                <IssueCard issue={issue} />
              </Box>
            );
          })}
        </Collapse>
      </CardContent>
    </Card>
  );
};

export default RepositoryCard;
