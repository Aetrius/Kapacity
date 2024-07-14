import * as React from 'react';
import {
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Divider,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography
} from '@mui/material';
import { DotsThree as DotsThreeIcon } from '@phosphor-icons/react/dist/ssr/DotsThree';
import { Storage } from './storage';

export interface StorageProps {
  storages: Storage[];
}

const headers = ['Name', 'Cluster', 'Memory Value', 'Memory Request', 'Memory Limit', 'CPU Value', 'CPU Request', 'CPU Limit'];

export function StorageTable({ storages }: StorageProps): React.JSX.Element {
  return (
    <Card>
      <CardHeader
        title="Storage"
      />
      <CardContent sx={{ pb: '8px' }}>
        <Box sx={{ minWidth: 800 }}>
          <Table>
            <TableHead>
              <TableRow>
                {headers.map((header) => (
                  <TableCell key={header}>
                    <Typography variant="subtitle2" fontWeight="bold">
                      {header}
                    </Typography>
                  </TableCell>
                ))}
                <TableCell />
              </TableRow>
            </TableHead>
            <TableBody>
              {storages.map((storage) => (
                <TableRow hover key={storage.id}>
                  <TableCell>
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <Typography noWrap variant="subtitle2">
                        {storage.title}
                      </Typography>
                    </Box>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{storage.cluster}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{storage.memoryValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{storage.memoryRequest}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{storage.memoryLimit}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{storage.cpuValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{storage.cpuRequest}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{storage.cpuLimit}</Typography>
                  </TableCell>
                  <TableCell>
                    <IconButton>
                      <DotsThreeIcon weight="bold" />
                    </IconButton>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </Box>
      </CardContent>
      <Divider />
      <CardActions>
        <Button color="secondary" endIcon={<DotsThreeIcon />} size="small">
          See all storage components
        </Button>
      </CardActions>
    </Card>
  );
}
