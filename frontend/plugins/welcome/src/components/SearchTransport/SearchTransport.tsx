import React, { useState, useEffect } from 'react';
import {
    Content,
    Header,
    Page,
    Link,
    pageTheme,
    ContentHeader,
} from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';

import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { Alert } from '@material-ui/lab';
import { EntUser } from '../../api/models/EntUser';
import { EntTransport } from '../../api/models/EntTransport';
import { EntHospital } from '../../api/models/EntHospital';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        table: {
            minWidth: 650,
        },
        margin: {
            margin: theme.spacing(3),
        },
    }),
);
const searchcheck = {
    ambulancesearchcheck: true,
    sendchcheck: true,
    receivecheck: true
}
export default function CarInspectionSearchPage() {
    const classes = useStyles();
    const api = new DefaultApi();
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบส่งตัวผู้ป่วย' };
    const [transports, setTransports] = useState([]);
    const [loading, setLoading] = useState(true);
    const [status, setStatus] = useState(false);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);
    const [users, setUsers] = useState<EntUser[]>([]);
    const [userid, setUser] = useState(Number);
    const [send, setSends] = useState<EntHospital[]>([]);
    const [receive, setReceivs] = useState<EntHospital[]>([]);


    

    const [sendsearch, setSendSearch] = useState(Number);
    const [receivesearch, setReceiveSearch] = useState(Number);
    const [ambulancesearch, setAmbulanceSearch] = useState(String);
    
    useEffect(() => {

       
        const getSends = async () => {
            const rs = await api.listHospital();
            setLoading(false);
            setSends(rs);
          };
          getSends();

        const getReceives = async () => {
            const rs = await api.listHospital();
            setLoading(false);
            setReceivs(rs);
          };
          getReceives();
        const getUsers = async () => {
            const us = await api.listUser();
            setLoading(false);
            setUsers(us);
          };
          getUsers();
          const checkJobPosition = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if (jobdata != "เจ้าหน้าที่รถพยาบาล") {
              localStorage.setItem("userdata", JSON.stringify(null));
              localStorage.setItem("jobpositiondata", JSON.stringify(null));
              history.pushState("", "", "./");
              window.location.reload(false);        
            }
            else {
              setUser(Number(localStorage.getItem("userdata")))
            }
          }
          checkJobPosition();

    }, [loading]);

    

    const SearchTransport = async () => {
        const apiUrl = `http://localhost:8080/api/v1/searchtransports?ambulance=${ambulancesearch}&send=${sendsearch}&receive=${receivesearch}`;
        const requestOptions = {
            method: 'GET',
        };
        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data.data)
                setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
                setAlertType("error");
                setTransports([]);
                if (data.data != null) {
                    if(data.data.length >= 1) {
                        setErrorMessege("พบข้อมูลที่ค้นหา");
                        setAlertType("success");
                        console.log(data.data)
                        setTransports(data.data);
                    }
                }

                setStatus(true);
            });

    }

    const SendhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setSendSearch(event.target.value as number);
    };
    const ReceivehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setReceiveSearch(event.target.value as number);
    };
    const AmbulanceSearchhandleChange = (event: any) => {
        setAmbulanceSearch(event.target.value as string);
    };
    


    return (
        <Page theme={pageTheme.service}>
            <Header
                title={`${profile.givenName}`}
            //subtitle="Some quick intro and links."
            ></Header>
            <Content>
                <ContentHeader title="การส่งตัวผู้ป่วย">
                {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
                    <Link component={RouterLink} to="/CreateTransport">
            <Button variant="contained" color="primary" style={{ backgroundColor: "#21b6ae" }}>
              กลับสู่หน้าเพิ่มการส่งตัวผู้ป่วย
            </Button>
          </Link>
                </ContentHeader>
                <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="send-label">โรงพยาบาลต้นทาง</InputLabel>
                <Select
                  labelId="send-label"
                  id="send"
                  value={sendsearch}
                  onChange={SendhandleChange}
                  style={{ width: 300 }}
                >
                  {send.map((item: EntHospital) => (
                    <MenuItem value={item.id}>{item.hospital}</MenuItem>
                  ))}
                </Select>
              </FormControl>
                
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="receive-label">โรงพยาบาลปลายทาง</InputLabel>
                <Select
                  labelId="receive-label"
                  id="receive"
                  value={receivesearch}
                  onChange={ReceivehandleChange}
                  style={{ width: 300 }}
                >
                  {receive.map((item: EntHospital) => (
                    <MenuItem value={item.id}>{item.hospital}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            
                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <TextField
                        id="search"
                        label="ค้นหาเลขทะเบียนรถ"
                        type="string"
                        size="medium"
                        value={ambulancesearch}
                        onChange={AmbulanceSearchhandleChange}
                        style={{ width: 400 }}
                    />
                </FormControl>

                <div className={classes.margin}>
                    <Button
                        onClick={() => {
                            SearchTransport();
                        }}
                        variant="contained"
                        color="primary"
                    >
                        ค้นหา
                            </Button>
                </div>

                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">เจ้าหน้าที่</TableCell>
                                <TableCell align="center">รถพยาบาล</TableCell>
                                <TableCell align="center">โรงพยาบาลต้นทาง</TableCell>
                                <TableCell align="center">โรงพยาบาลปลายทาง</TableCell>
                                <TableCell align="center">อาการผู้ป่วย</TableCell>
                                <TableCell align="center">การแพ้ยา</TableCell>
                                <TableCell align="center">หมายเหตุเพิ่มเติม</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {transports.map((item: any) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges?.User?.name}</TableCell>
                                    <TableCell align="center">{item.edges?.Ambulance?.carregistration}</TableCell>
                                    <TableCell align="center">{item.edges?.Send?.hospital}</TableCell>
                                    <TableCell align="center">{item.edges?.Receive?.hospital}</TableCell>
                                    <TableCell align="center">{item.symptom}</TableCell>
                                    <TableCell align="center">{item.drugallergy}</TableCell>
                                    <TableCell align="center">{item.note}</TableCell>
                                    
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Content>
        </Page>
    );
}