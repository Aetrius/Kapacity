import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import IconButton from '@mui/material/IconButton';
import ListItem from '@mui/material/ListItem';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemText from '@mui/material/ListItemText';
import Typography from '@mui/material/Typography';
import { Application } from './application';
import { SquaresFour as SquaresFourIcon } from '@phosphor-icons/react/dist/ssr/SquaresFour';
import { DotsThree as DotsThreeIcon } from '@phosphor-icons/react/dist/ssr/DotsThree';

export function ApplicationItem({ application }: { application: Application }): React.JSX.Element {
  return (
    <ListItem disableGutters sx={{ display: 'flex', alignItems: 'center' }}>
      <ListItemAvatar>
        <SquaresFourIcon size={32} color="#30c41c" />
      </ListItemAvatar>
      <ListItemText sx={{ flex: 1 }}>
        <Typography noWrap variant="subtitle2">
          {application.title}
        </Typography>
      </ListItemText>
      <ListItemText sx={{ flex: 1 }}>
        <Typography sx={{ whiteSpace: 'nowrap' }} variant="body2">
          {application.memoryValue}
        </Typography>
      </ListItemText>
      <ListItemText sx={{ flex: 1 }}>
        <Typography sx={{ whiteSpace: 'nowrap' }} variant="body2">
          {application.memoryRequest}
        </Typography>
      </ListItemText>
      <ListItemText sx={{ flex: 1 }}>
        <Typography sx={{ whiteSpace: 'nowrap' }} variant="body2">
          {application.memoryLimit}
        </Typography>
      </ListItemText>
      <ListItemText sx={{ flex: 1 }}>
        <Typography sx={{ whiteSpace: 'nowrap' }} variant="body2">
          {application.cpuValue}
        </Typography>
      </ListItemText>
      <ListItemText sx={{ flex: 1 }}>
        <Typography sx={{ whiteSpace: 'nowrap' }} variant="body2">
          {application.cpuRequest}
        </Typography>
      </ListItemText>
      <ListItemText sx={{ flex: 1 }}>
        <Typography sx={{ whiteSpace: 'nowrap' }} variant="body2">
          {application.cpuLimit}
        </Typography>
      </ListItemText>
      <IconButton>
        <DotsThreeIcon weight="bold" />
      </IconButton>
    </ListItem>
  );
}
