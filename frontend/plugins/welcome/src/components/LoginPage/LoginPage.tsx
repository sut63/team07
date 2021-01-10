import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    media: {
      height: 140,
    },
  }),
);

export default function Login(props: any) {
  const classes = useStyles();
  const profile = { givenName: 'to Software Analysis 63' };
  const api = new DefaultApi();

  const [users, setUsers] = useState<EntUser[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);

  const [email, setEmail] = useState(String);
  const [password, setPassword] = useState(String);

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getUsers = async () => {
      const res: any = await api.listUser();
      setLoading(false);
      setUsers(res);
    }

    getUsers();
  }, [loading]);

  const EmailthandleChange = (event: any) => {
    setEmail(event.target.value as string);
  };

  const PasswordthandleChange = (event: any) => {
    setPassword(event.target.value as string);
  };

  const LoginChecker = async () => {
    users.map((item: any) => {
      console.log(item.email);
      if ((item.email == email) && (item.password == password)) {
        setStatus(true);
        setAlert(true);
        localStorage.setItem("userdata",JSON.stringify(item.id));
        localStorage.setItem("jobpositiondata",JSON.stringify(item.edges.jobposition.positionName))
        if (item.edges.jobposition.positionName == "เจ้าหน้าที่ตรวจสภาพรถ"){
          history.pushState("","","/carinspection");
        }
        else if (item.edges.jobposition.positionName == "เจ้าหน้าที่รถพยาบาล"){
          history.pushState("","","/maindriver");
        }
        else if (item.edges.jobposition.positionName == "เจ้าหน้าที่โอเปอร์เรเตอร์"){
          history.pushState("","","/carservicemain");
        }
        else if (item.edges.jobposition.positionName == "เจ้าหน้าที่ซ่อมบำรุงรถ"){
          history.pushState("","","/carrepairmain");
        }


        window.location.reload(false);

        const timer = setTimeout(() => {
        }, 1000);
      }
      else {
        setStatus(true);
        setAlert(false);
      }
    })
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ยินดีต้อนรับสู่ ระบบรถพยาบาล`}
        subtitle="บริการ คุ ณ ภ า พ"
      ></Header>
      <Content>
        <ContentHeader title="Login">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="email"
                  label="Email"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={email}
                  onChange={EmailthandleChange}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="password"
                  label="Password"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={password}
                  onChange={PasswordthandleChange}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div>

            <div className={classes.margin}>
              <Button
                style={{ width: 400, backgroundColor: "#34eb77" }}
                onClick={() => {
                  LoginChecker();
                }}
                variant="contained"
                color="primary"
              >
                Enter
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
