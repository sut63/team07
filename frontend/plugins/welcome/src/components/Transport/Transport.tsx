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

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import SaveIcon from '@material-ui/icons/Save';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { EntAmbulance } from '../../api/models/EntAmbulance';
import { EntSend } from '../../api/models/EntSend';
import { EntReceive } from '../../api/models/EntReceive';
import { EntUser } from '../../api/models/EntUser';
import ComponanceTransportTable from '../TransportTable';
import { EntTransport } from '../../api';
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
  const [ambulanceid, setAmbulance] = useState(Number);
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
      const checkJobPosition = async () => {
        const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
        setLoading(false);
        if (jobdata != "เจ้าหน้าที่รถพยาบาล" ) {
          localStorage.setItem("userdata",JSON.stringify(null));
          localStorage.setItem("jobpositiondata",JSON.stringify(null));
          history.pushState("","","./");
          window.location.reload(false);        
        }
        else{
            setUser(Number(localStorage.getItem("userdata")))
        }
      }
    checkJobPosition();
 
  }, [loading]);

  console.log(userid)
  

  const SendhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setsend(event.target.value as number);
  };

  const ReceivehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setreceive(event.target.value as number);
  };

  const AmbulancehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAmbulance(event.target.value as number);
  };

  const UserthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };


  const CreateTransport = async () => {
    if ((sendid != 0) && (receiveid != 0) && (ambulanceid != 0)){
    const transports = {
      sendID: sendid,
      receiveID: receiveid,
      ambulanceID: ambulanceid,
      userID: userid,
   };
    console.log(transports);
    const res: any = await api.createTransport({ transport: transports });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      window.location.reload(false);
    }
  }
     else {
      setAlert(false);
      setStatus(true);
    }
  };
  

  return (
 <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
         >
      </Header>
      <Content>
      <ContentHeader title="กรอกข้อมูลส่งตัวผู้ป่วย">
                    {status ? (
                        <div>
                            {alert ? (
                                <Alert severity="success">
                                    บันทึกสำเร็จ
                                </Alert>
                            ) : (
                                <Alert severity="warning" style={{marginTop: 20}}>
                                    กรุณากรอกข้อมูลให้ครบถ้วน
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
            </div>      
            
            
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
                <div className={classes.paper}><strong>ชื่อเจ้าหน้าที่รถพยาบาล</strong></div>
                <TextField
                                    id="user"
                                    type="string"
                                    size="medium"
                                    value={users.filter((filter:EntUser) => filter.id == userid).map((item:EntUser) => `${item.name} (${item.email})`)}
                                    style={{ width: 400 }}
                                />
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
                  value={ambulanceid}
                  onChange={AmbulancehandleChange}
                  style={{ width: 400 }}
                >
                  {ambulances.filter((filter:any) => filter.edges.Hasstatus.resultName == "พร้อมใช้งาน").map((item: any) => (
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
                to="/maindriver"
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
