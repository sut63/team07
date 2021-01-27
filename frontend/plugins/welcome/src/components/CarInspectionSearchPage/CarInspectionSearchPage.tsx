import React, { useState, useEffect } from 'react';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
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

import { EntCarInspection } from '../../api/models/EntCarInspection';
import { EntInspectionResult } from '../../api/models/EntInspectionResult';

import SearchIcon from '@material-ui/icons/Search';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';

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
    resultsearchcheck: true
}

export default function CarInspectionSearchPage() {
    const classes = useStyles();
    const api = new DefaultApi();
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบตรวจสภาพรถ' };
    const [carinspections, setCarInspections] = useState<EntCarInspection[]>([]);
    const [loading, setLoading] = useState(true);
    const [status, setStatus] = useState(false);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [inspectionresults, setInspectionResults] = useState<EntInspectionResult[]>([]);

    const [inspectionresultsearch, setInspectionResult] = useState(Number);
    const [usersearch, setUserSearch] = useState(String);
    const [ambulancesearch, setAmbulanceSearch] = useState(String);

    useEffect(() => {
        const getInspectionResults = async () => {
            const res = await api.listInspectionresult();
            setLoading(false);
            setInspectionResults(res);
        };
        getInspectionResults();

        const checkJobPosition = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if (jobdata != "เจ้าหน้าที่ตรวจสภาพรถ") {
                localStorage.setItem("userdata", JSON.stringify(null));
                localStorage.setItem("jobpositiondata", JSON.stringify(null));
                history.pushState("", "", "./");
                window.location.reload(false);
            }
        }
        checkJobPosition();


    }, [loading]);

    const SearchCarInspection = async () => {
        const res = await api.listCarinspection();
        const usersearch = UserSearch(res);
        const ambulancesearch = AmbulanceSearch(usersearch);
        const resultsearch = InspectionResultSearch(ambulancesearch);
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setCarInspections([]);
        if(resultsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setCarInspections(resultsearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.ambulancesearchcheck = true;
        searchcheck.usersearchcheck = true;
        searchcheck.resultsearchcheck = true;
    }

    const AmbulanceSearch = (res: any) => {
        const data = res.filter((filter: EntCarInspection) => filter.edges?.ambulance?.carregistration?.includes(ambulancesearch))
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
        const data = res.filter((filter: EntCarInspection) => filter.edges?.user?.name?.includes(usersearch))
        if (data.length != 0  && usersearch != "") {
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

    const InspectionResultSearch = (res: any) => {
        const data = res.filter((filter: EntCarInspection) => filter.edges?.inspectionresult?.id == inspectionresultsearch)
        if (data.length != 0) {
            return data;
        }
        else {
            searchcheck.resultsearchcheck = false;
            if(inspectionresultsearch == 0){
                return res;
            }
            else{
                return data;
            }
        }
    }


    const InspectionResulthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setInspectionResult(event.target.value as number);
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
            //subtitle="Some quick intro and links."
            ></Header>
            <Content>
                <ContentHeader title="ค้นหาข้อมูลการตรวจสภาพรถ">
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
                        label="ค้นหาผู้บันทึกด้วย keyword"
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
                        label="ค้นหาเลขทะเบียนรถด้วย keyword"
                        type="string"
                        size="medium"
                        value={ambulancesearch}
                        onChange={AmbulanceSearchhandleChange}
                        style={{ width: 300 }}
                    />
                </FormControl>

                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <InputLabel id="result-label">ค้นหาผลการตรวจสภาพรถ</InputLabel>
                    <Select
                        labelId="result-label"
                        id="result"
                        value={inspectionresultsearch}
                        onChange={InspectionResulthandleChange}
                        style={{ width: 200 }}

                    >   <MenuItem value={0}>-</MenuItem>
                        {inspectionresults.filter((filter: EntInspectionResult) => filter.edges?.jobposition?.positionName == "เจ้าหน้าที่ตรวจสภาพรถ").map((item: EntInspectionResult) => (
                            <MenuItem value={item.id}>{item.resultName}</MenuItem>
                        ))}
                    </Select>
                </FormControl>

                <div className={classes.margin}>
                    <Button
                        onClick={() => {
                            SearchCarInspection();
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
                            history.pushState("", "", "./carinspection");
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
                                <TableCell align="center">ศูนย์ล้อ (เมตร)</TableCell>
                                <TableCell align="center">ระดับเสียง (เดซิเบล)</TableCell>
                                <TableCell align="center">ควันดำ (เปอร์เซ็นต์)</TableCell>
                                <TableCell align="center">ผลการตรวจสภาพรถ</TableCell>
                                <TableCell align="center">วัน/เดือน/ปี เวลา</TableCell>
                                <TableCell align="center">หมายเหตุ</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {carinspections.map((item: any) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges.user.name}</TableCell>
                                    <TableCell align="center">{item.edges.ambulance.carregistration}</TableCell>
                                    <TableCell align="center">{item.wheelCenter}</TableCell>
                                    <TableCell align="center">{item.soundLevel}</TableCell>
                                    <TableCell align="center">{item.blacksmoke}</TableCell>
                                    <TableCell align="center">{item.edges.inspectionresult.resultName}</TableCell>
                                    <TableCell align="center">{moment(item.datetime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
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