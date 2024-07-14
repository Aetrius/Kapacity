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
import { Cluster } from './cluster';

export interface ClusterProps {
  clusters: Cluster[];
}

const headers = ['Name', 'Memory Value', 'Memory Limit', 'CPU Value', 'CPU Limit', 'diskValue', 'diskSize', 'location', 'networkSpeed'];

export function ClusterTable({ clusters }: ClusterProps): React.JSX.Element {
  return (
    <Card>
      <CardHeader
        title="Clusters"
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
              {clusters.map((cluster) => (
                <TableRow hover key={cluster.id}>
                  <TableCell>
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <Typography noWrap variant="subtitle2">
                        {cluster.name}
                      </Typography>
                    </Box>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{cluster.memoryValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{cluster.memoryLimit}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{cluster.cpuValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{cluster.cpuLimit}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{cluster.diskValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{cluster.diskSize}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{cluster.location}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{cluster.networkSpeed}</Typography>
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
          See all clusters
        </Button>
      </CardActions>
    </Card>
  );
}
