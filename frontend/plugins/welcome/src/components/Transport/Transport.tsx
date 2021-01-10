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
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { EntAmbulance } from '../../api/models/EntAmbulance';
import { EntSend } from '../../api/models/EntSend';
import { EntReceive } from '../../api/models/EntReceive';
import { EntUser } from '../../api/models/EntUser';
import ComponanceTransportTable from '../TransportTable';
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
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    formControl: {
      width: 400,
    },
  }),
);


export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบขนส่งผู้ป่วย' };
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  //เก็บข้อมูลที่จะดึงมา
  const [sends, setSends] = useState<EntSend[]>([]);
  const [receives, setReceives] = useState<EntReceive[]>([]);
  const [ambulances, setAmbulances] = useState<EntAmbulance[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);

  const [loading, setLoading] = useState(true);

  const [sendid, setsend] = useState(Number);
  const [receiveid, setreceive] = useState(Number);
  const [ambulanceid, setambulance] = useState(Number);
  const [userid, setUser] = useState(Number);

  useEffect(() => {

    const getSends = async () => {
 
      const sd = await api.listSend();
      setLoading(false);
      setSends(sd);
    };
    getSends();
 
    const getUsers = async () => {
 
    const us = await api.listUser();
      setLoading(false);
      setUsers(us);
    };
    getUsers();
 
    const getReceives = async () => {
 
     const rs = await api.listReceive();
       setLoading(false);
       setReceives(rs);
     };
     getReceives();

     const getAmbulances = async () => {
 
      const am = await api.listAmbulance({ limit: 10, offset: 0 });
        setLoading(false);
        setAmbulances(am);
      };
      getAmbulances();
 
  }, [loading]);

  console.log(userid)
  

  const SendhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setsend(event.target.value as number);
  };

  const ReceivehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setreceive(event.target.value as number);
  };

  const AmbulancehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setambulance(event.target.value as number);
  };

  const UserthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const CreateTransport = async () => {
    const transports = {
      sendID: sendid,
      recieveID: receiveid,
      ambulanceID: ambulanceid,
      userID: userid,

    };
    console.log(transports);
    const res: any = await api.createTransport({ transport : transports });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      window.location.reload(false);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };
  

  return (
 <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      //subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลรถเข้าสู่ระบบ">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    กรุณากรอกข้อมูลอีกครั้ง
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>โรงพยาบาลที่จะรับ</strong></div>
              <InputLabel id="receive-label"></InputLabel>
              <Select
                labelId="receive-label"
                id="receive"
                value={receiveid}
                onChange={ReceivehandleChange}
                style={{ width: 400 }}
              >
                {receives.map((item: EntReceive) => (
                  <MenuItem value={item.id}>{item.receivename}</MenuItem>
                ))}
              </Select>
            </FormControl>

            <div></div>
            
            
            <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>โรงพยาบาลที่จะส่ง</strong></div>
              <InputLabel id="send-label"></InputLabel>
              <Select
                labelId="send-label"
                id="send"
                value={sendid}
                onChange={SendhandleChange}
                style={{ width: 400 }}
              >
                {sends.map((item: EntSend) => (
                  <MenuItem value={item.id}>{item.sendname}</MenuItem>
                ))}
              </Select>
            </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>เจ้าหน้าที่รถพยาบาล</strong></div>
                <InputLabel id="user-label"></InputLabel>
                <Select
                  labelId="user-label"
                  id="user"
                  value={userid}
                  onChange={UserthandleChange}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
                ))}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>รถพยาบาล</strong></div>
                <InputLabel id="ambulance-label"></InputLabel>
                <Select
                  labelId="ambulance-label"
                  id="ambulance"
                  value={ambulances}
                  onChange={AmbulancehandleChange}
                  style={{ width: 400 }}
                >
                  {ambulances.map((item: EntAmbulance) => (
                  <MenuItem value={item.id}>{item.carregistration}</MenuItem>
                ))}
                </Select>
              </FormControl>
            </div>

            
            
                <Typography align ="right"></Typography>

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateTransport();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/mainambulance"
                variant="contained"
              >
                กลับ
             </Button>
            </div>
          </form>
        </div>
        
        <ComponanceTransportTable></ComponanceTransportTable>
      </Content>
    </Page>
  );
}
