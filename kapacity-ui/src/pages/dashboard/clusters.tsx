import * as React from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Unstable_Grid2';
import { Cpu } from '@phosphor-icons/react/dist/ssr/Cpu';
import { HardDrives } from '@phosphor-icons/react/dist/ssr/HardDrives';
import { Memory } from '@phosphor-icons/react/dist/ssr/Memory';
import { Helmet } from 'react-helmet-async';

import type { Metadata } from '@/types/metadata';
import { config } from '@/config';
import { Summary } from '@/components/dashboard/overview/summary';
import { SummaryPercent } from '@/components/dashboard/overview/summaryPerecent';
import { Cube as CubeIcon } from '@phosphor-icons/react/dist/ssr/Cube';
import { Dropdown } from '@/components/core/dropdown/dropdown';
import { DropdownTrigger } from '@/components/core/dropdown/dropdown-trigger';
import { DropdownPopover } from '@/components/core/dropdown/dropdown-popover';
import { MenuList, MenuItem } from '@mui/material';
import { ClusterTable } from '@/components/dashboard/overview/clusterTable';
import { Cluster } from '@/components/dashboard/overview/cluster';

const metadata = { title: `Dashboard | Clusters | ${config.site.name}` } satisfies Metadata;


const clusters: Cluster[] = [
  {
    id: '1',
    name: 'Cluster One',
    icon: 'icon1.png',
    memoryValue: '12Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuLimit: '5000m',
    location: 'dc1',
    diskValue: '50Gib',
    diskSize: '100Gib',
    networkSpeed: '100mbps',
  },
  {
    id: '11',
    name: 'Cluster Two',
    icon: 'icon1.png',
    memoryValue: '12Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuLimit: '5000m',
    location: 'dc1',
    diskValue: '50Gib',
    diskSize: '100Gib',
    networkSpeed: '100mbps',
  },
  {
    id: '12',
    name: 'Cluster Three',
    icon: 'icon1.png',
    memoryValue: '12Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuLimit: '5000m',
    location: 'dc1',
    diskValue: '50Gib',
    diskSize: '100Gib',
    networkSpeed: '100mbps',
  },
  {
    id: '13',
    name: 'Cluster Four',
    icon: 'icon1.png',
    memoryValue: '12Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuLimit: '5000m',
    location: 'dc1',
    diskValue: '50Gib',
    diskSize: '100Gib',
    networkSpeed: '100mbps',
  },
];

export function Clusters(): React.JSX.Element {
  return (
    <React.Fragment>
      <Helmet>
        <title>{metadata.title}</title>
      </Helmet>
      <Box
        sx={{
          maxWidth: 'var(--Content-maxWidth)',
          m: 'var(--Content-margin)',
          p: 'var(--Content-padding)',
          width: 'var(--Content-width)',
        }}
      >
        <Stack spacing={4}>
          <Stack direction={{ xs: 'column', sm: 'row' }} spacing={3} sx={{ alignItems: 'flex-start' }}>
            <Box sx={{ flex: '1 1 auto' }}>
              <Typography variant="h4">Clusters</Typography>
            </Box>
            <Box>
            <Dropdown>
              <DropdownTrigger>
              <Button variant="contained">Environment: All</Button>
              </DropdownTrigger>
              <DropdownPopover>
                <MenuList>
                  <MenuItem>PROD</MenuItem>
                  <MenuItem>STAGE</MenuItem>
                  <MenuItem>QA</MenuItem>
                </MenuList>
              </DropdownPopover>
            </Dropdown>
            </Box>
          </Stack>
          <Grid container spacing={4}>
            <Grid md={12} xs={12}>
              <ClusterTable clusters={clusters}/>
            </Grid>
          </Grid>
        </Stack>
      </Box>
    </React.Fragment>
  );
}
