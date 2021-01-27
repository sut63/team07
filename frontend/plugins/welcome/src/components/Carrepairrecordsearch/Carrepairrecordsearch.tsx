import React, { useState, useEffect, Fragment } from 'react';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
    ItemCard,
} from '@backstage/core';
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

import moment from 'moment';

import { EntCarRepairrecord } from '../../api/models/EntCarRepairrecord';
import { EntRepairing } from '../../api/models/EntRepairing';

import SearchIcon from '@material-ui/icons/Search';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import { EntCarInspection, EntUser } from '../../api';

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
    usersearchcheck: true,
    ambulancesearchcheck: true,
    repairingsearchcheck: true
}

export default function CarRepairrecordSearch() {
    const classes = useStyles();
    const api = new DefaultApi();
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบซ่อมบำรุง' };
    const [carrepairrecod, setCarrepairrecords] = useState<EntCarRepairrecord[]>([]);
    const [loading, setLoading] = useState(true);
    const [status, setStatus] = useState(false);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);
    const [carinspections, setCarInspections] = useState<EntCarInspection[]>([]);

    const [repairings, setRepairings] = useState<EntRepairing[]>([]);

    const [usersearch, setUserSearch] = useState(String);
    const [ambulancesearch, setAmbulanceSearch] = useState(String);
    const [repairingsearch, setRepairingSearch] = useState(Number);

    useEffect(() => {
        const getRepairings = async () => {
            const res = await api.listRepairing({offset:0 , limit:100});
            setLoading(false);
            setRepairings(res);
        };
        getRepairings();

        const getCarinspections = async () => {
            const res2 = await api.listCarinspection();
            setLoading(false);
            setCarInspections(res2);
          };
          getCarinspections();

        const checkJobPosition = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if (jobdata != "เจ้าหน้าที่ซ่อมบำรุงรถ") {
                localStorage.setItem("userdata", JSON.stringify(null));
                localStorage.setItem("jobpositiondata", JSON.stringify(null));
                history.pushState("", "", "./");
                window.location.reload(false);
            }
        }
        checkJobPosition();

    }, [loading]);

    const SearchCarRepairrecord = async () => {
        const res = await api.listCarrepairrecord();
        const usersearch = UserSearch(res);
        const ambulancesearch = AmbulanceSearch(usersearch);
        const repairingsearch = RepairingSearch(ambulancesearch);

        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setCarrepairrecords([]);
        if(repairingsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) => {
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setCarrepairrecords(repairingsearch);
                }
            })
        }
        
        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.ambulancesearchcheck = true;
        searchcheck.repairingsearchcheck = true;
        searchcheck.usersearchcheck = true;
    }

    const AmbulanceSearch = (res: any) => {
        const datafilter = carinspections.filter((item:EntCarInspection) => item.edges?.ambulance?.carregistration?.includes(ambulancesearch));
        const data = res.filter((filter:any) => filter.edges?.carinspection?.id == datafilter.map((inspection:EntCarInspection) => inspection.id))
        console.log(data)
        if (data.length != 0 && ambulancesearch != "") {
            return data;
        }
        else {
            searchcheck.ambulancesearchcheck = false;
            if(ambulancesearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const UserSearch = (res: any) => {
        const data = res.filter((filter: EntCarRepairrecord) => filter.edges?.user?.name?.includes(usersearch))
        if (data.length != 0 && usersearch != "") {
            return data;
        }
        else {
            searchcheck.usersearchcheck = false;
            if(usersearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const RepairingSearch = (res: any) => {
        const data = res.filter((filter: EntCarRepairrecord) => filter.edges?.keeper?.id == repairingsearch)
        if (data.length != 0) {
            return data;
        }
        else {
            searchcheck.repairingsearchcheck = false;
            if(data.length != 0) {
                return data;
            }
            else {
                searchcheck.repairingsearchcheck = false;
                if(repairingsearch == 0){
                    return res;
                }
                else{
                    return data;
                }
            }
        }
    }

    const RepairinghandleChange = (event: React.ChangeEvent<{value: unknown}>) => {
        setRepairingSearch(event.target.value as number);
    };

    const UserSearchhandleChange = (event: any) => {
        setUserSearch(event.target.value as string);
    };

    const AmbulanceSearchhandleChange = (event: any) => {
        setAmbulanceSearch(event.target.value as string);
    };

    return (
        <Page theme={pageTheme.service}>
            <Header
                title={`${profile.givenName}`}
            ></Header>
            <Content>
                <ContentHeader title="ค้นหาข้อมูลซ่อมบำรุง">
                    {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
                </ContentHeader>
                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <TextField
                        id="search"
                        label="ค้นหาผู้บันทึก"
                        type="string"
                        size="medium"
                        value={usersearch}
                        onChange={UserSearchhandleChange}
                        style={{ width: 200 }}
                    />
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
                        style={{ width: 300}}
                    />
                </FormControl>

                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <InputLabel id="repair-label">ค้นหาส่วนที่ซ่อม</InputLabel>
                    <Select
                        labelId="repair-label"
                        id="repair"
                        value={repairingsearch}
                        onChange={RepairinghandleChange}
                        style={{ width: 200 }}
                    >
                        <MenuItem value={0}>-</MenuItem>
                        {repairings.map((item: EntRepairing) =>(
                            <MenuItem value={item.id}>{item.repairpart}</MenuItem>
                        ))}
                    </Select>
                </FormControl>

                <div className={classes.margin}>
                    <Button
                        onClick={() => {
                            SearchCarRepairrecord();
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 180, backgroundColor: "#365391" }}
                        startIcon={<SearchIcon />}
                    >
                        ค้นหา
                    </Button>
                    <Button
                        onClick={() => {
                            history.pushState("", "", "./carrepairmain");
                            window.location.reload(false);
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 180, marginLeft: 40, backgroundColor: "#a31f1f" }}
                        startIcon={<ArrowBackIcon />}
                    >
                        กลับ
                    </Button>
                </div>

                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">เจ้าหน้าที่</TableCell>
                                <TableCell align="center">รถพยาบาล</TableCell>
                                <TableCell align="center">ส่วนที่ซ่อม</TableCell>
                                <TableCell align="center">หมายเหตุส่วนที่ซ่อม</TableCell>
                                <TableCell align="center">เงินที่ใช้สำหรับการซ่อม</TableCell>
                                <TableCell align="center">ความคิดเห็นจากช่าง</TableCell>
                                <TableCell align="center">วัน/เดือน/ปี เวลา</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {carrepairrecod.map((item: EntCarRepairrecord) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges?.user?.name}</TableCell>
                                    {carinspections.filter((item2: EntCarInspection) => item2.id == item.edges?.carinspection?.id).map((item3: EntCarInspection) => (
                                        <TableCell align="center">{item3.edges?.ambulance?.carregistration}</TableCell>
                                    ))}
                                    <TableCell align="center">{item.edges?.keeper?.repairpart}</TableCell>
                                    <TableCell align="center">{item.partrepair}</TableCell>
                                    <TableCell align="center">{item.price}</TableCell>
                                    <TableCell align="center">{item.techniciancomment}</TableCell>
                                    <TableCell align="center">{moment(item.datetime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>        
            </Content>
        </Page>
    );
}