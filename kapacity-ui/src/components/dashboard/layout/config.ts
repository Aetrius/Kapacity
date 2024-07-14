import type { NavItemConfig } from '@/types/nav';
import { paths } from '@/paths';

// NOTE: We did not use React Components for Icons, because
//  you may one to get the config from the server.

// NOTE: First level elements are groups.

export interface LayoutConfig {
  navItems: NavItemConfig[];
}

export const layoutConfig = {
  navItems: [
    {
      key: 'dashboards',
      title: 'Dashboards',
      items: [
        { key: 'overview', title: 'Overview', href: paths.dashboard.overview, icon: 'house' },
        { key: 'clusters', title: 'Clusters', href: paths.dashboard.clusters, icon: 'chart-pie' },
        { key: 'applications', title: 'Applications', href: paths.dashboard.applications, icon: 'cube' },
        { key: 'nodes', title: 'Nodes', href: paths.dashboard.nodes, icon: 'upload' },
        // { key: 'storage', title: 'Storage', href: paths.dashboard.storage, icon: 'disk' },
      ],
    },
    // {
    //   key: 'administration',
    //   title: 'Administration',
    //   items: [
    //     {
    //       key: 'settings',
    //       title: 'Settings',
    //       href: paths.dashboard.settings.account,
    //       icon: 'gear',
    //       matcher: { type: 'startsWith', href: '/dashboard/settings' },
    //     },
    //     { key: 'tasks', title: 'Tasks', href: paths.dashboard.tasks, icon: 'kanban' },
    //   ],
    // }
    
  ],
} satisfies LayoutConfig;
