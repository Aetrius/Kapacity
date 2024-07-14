import React from 'react';
import {
  List,
  ListItem,
  ListItemText,
  Card,
  CardContent,
  Typography,
  Divider,
  Box,
  Stack
} from '@mui/material';
import { DetailItem } from './detailItem';

interface DetailListProps {
  items: DetailItem[];
  title: string;
}

const DetailList: React.FC<DetailListProps> = ({ title, items }) => {
  return (
    <Card>
      <CardContent>
        <Typography color="text.secondary" variant="body1">
          {title}
        </Typography>
        <List>
          {items.map((item, index) => (
            <ListItem key={index}>
              <ListItemText
                primary={item.primary}
                secondary={item.secondary}
              />
            </ListItem>
          ))}
        </List>
      </CardContent>
      <Divider />
      <CardContent>
          hello moto
      </CardContent>
    </Card>
  );
};

export default DetailList;
