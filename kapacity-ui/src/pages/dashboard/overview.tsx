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

const metadata = { title: `Dashboard | Overview | ${config.site.name}` } satisfies Metadata;

export function Page(): React.JSX.Element {
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
              <Typography variant="h4">Overview</Typography>
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
          <Grid container spacing={2}>
            <Grid md={4} xs={12}>
              <SummaryPercent amount={31} diff={15} icon={Cpu} title="Compute Capacity" trend="up" />
            </Grid>
            <Grid md={4} xs={12}>
              <SummaryPercent amount={40} diff={5} icon={Memory} title="Memory Capacity" trend="up" />
            </Grid>
            <Grid md={4} xs={12}>
              <SummaryPercent amount={21} diff={12} icon={HardDrives} title="Storage Capacity" trend="up" />
            </Grid>
            <Grid md={4} xs={12}>
              <Summary amount={1200} total={12500} icon={CubeIcon} title="Pod Count" trend="up" />
            </Grid>
            <Grid md={4} xs={12}>
            <Summary amount={120} total={5000} icon={CubeIcon} title="Node Count" trend="up" />
            </Grid>
            <Grid md={4} xs={12}>
              <Summary amount={20000} total={300000} icon={CubeIcon} title="Container Count" trend="up" />
            </Grid>
            <Grid md={4} xs={12}>
              <Summary amount={25} total={110} icon={CubeIcon} title="Avg Pods Per Node" trend="up" />
            </Grid>
            
          </Grid>
        </Stack>
      </Box>
    </React.Fragment>
  );
}
