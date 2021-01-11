import React, { useState, useEffect } from 'react';
import PermIdentityIcon from '@material-ui/icons/PermIdentity';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import MeetingRoomIcon from '@material-ui/icons/MeetingRoom';
import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

import { EntUser } from 'plugin-welcome/src/api/models/EntUser';
import { DefaultApi } from 'plugin-welcome/src/api/apis';

export const AppSidebar = () => {

  const api = new DefaultApi();
  const [userid, setUser] = useState(Number);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
    };
    getUsers();
    const data = localStorage.getItem("userdata");
    if (data) {
      setUser(Number(JSON.parse(data)));
      setLoading(false);
    }
    
  }, [loading]);

  return (

    <Sidebar>
      <SidebarDivider />
      {/* Global nav, not org-specific */}
      {(userid) ?
        users.filter((filter:EntUser) => filter.id == userid).map((item:EntUser) => 
          <SidebarItem icon={PermIdentityIcon} text={item.name} />
        )
        :
        null
      }
      {/* End global nav */}
      <SidebarDivider />
      <SidebarSpace />
      <SidebarDivider />
      <SidebarThemeToggle />
      {(userid) ?
        <SidebarItem icon={MeetingRoomIcon} to="./" text="ออกจากระบบ"
          onClick={() => {
            localStorage.setItem("userdata", JSON.stringify(null));
            localStorage.setItem("jobpositiondata", JSON.stringify(null));
            history.pushState("", "", "./");
            window.location.reload(false);
          }} />
        :
        null
      }
      <SidebarPinButton />
    </Sidebar>
  )
};