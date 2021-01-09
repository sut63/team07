import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import BuildIcon from '@material-ui/icons/Build';
import SignOut from '@material-ui/icons/Settings';
import AirportShuttleIcon from '@material-ui/icons/AirportShuttle';
import AssignmentIcon from '@material-ui/icons/Assignment';
import AccessibleForwardIcon from '@material-ui/icons/AccessibleForward';
import AirlineSeatFlatAngledIcon from '@material-ui/icons/AirlineSeatFlatAngled';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';


import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    <SidebarItem
      icon={AirportShuttleIcon}
      to="Mainambulance"
      text="รถเข็ญใหม่"
    />
    <SidebarItem
      icon={AssignmentIcon}
      to="Carcheckinout"
      text="รถเข้าออก"
    />
    <SidebarItem
      icon={ShoppingCartIcon}
      to="CarInspection"
      text="ตรวจรถเข็ญ"
    />
    <SidebarItem
      icon={AccessibleForwardIcon}
      to="Carservicemain"
      text="บริการรถเข็ญ"
    />
    <SidebarItem
      icon={AirlineSeatFlatAngledIcon}
      to="watch_video" //ใส่ตัวแปรของplugin
      text="บีม"
    />
    <SidebarItem
      icon={BuildIcon}
      to="watch_video"  //ใส่ตัวแปรของplugin
      text="ฟง"
    />

    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to="sign_out"
      text="Sign Out"
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);
