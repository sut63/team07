import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntCarservice } from '../../api/models/EntCarservice';
import { EntUser } from '../../api/models/EntUser';
import moment from 'moment';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [carservices, setCarservices] = useState<EntCarservice[]>([]);
 const [loading, setLoading] = useState(true);
 const [users, setUsers] = useState<EntUser[]>([]);
 const [userid, setUser] = useState(Number);
 
 useEffect(() => {
   const getCarservices = async () => {
     const res = await api.listCarservice({ limit: 10, offset: 0 });
     setLoading(false);
     setCarservices(res);
   };
   getCarservices();

   const checkJobPosition = async () => {
    const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
    setLoading(false);
    if (jobdata != "เจ้าหน้าที่โอเปอร์เรเตอร์" ) {
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
 
 const deleteCarservices = async (id: number) => {
  const res = await api.deleteCarservice({ id: id });
  setLoading(true);
};

 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">เลขที่</TableCell>
           <TableCell align="center">ชื่อผู้ใช้บริการ</TableCell>
           <TableCell align="center">อายุผู้ใช้บริการ</TableCell>
           <TableCell align="center">ที่อยู่</TableCell>
           <TableCell align="center">การใช้บริการ</TableCell>
           <TableCell align="center">ความเร่งด่วน</TableCell>
           <TableCell align="center">ระยะทาง</TableCell>
           <TableCell align="center">ผู้บันทึก</TableCell>
           <TableCell align="center">วัน-เวลา</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {carservices.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.customer}</TableCell>
             <TableCell align="center">{item.age}</TableCell>
             <TableCell align="center">{item.location}</TableCell>
             <TableCell align="center">{item.information}</TableCell>
             <TableCell align="center">{item.edges?.urgentid?.urgent}</TableCell>
             <TableCell align="center">{item.edges?.disid?.distance}</TableCell>
             <TableCell align="center">{item.edges?.userid?.name}</TableCell>
             <TableCell align="center">{moment(item.datetime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteCarservices(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 ลบบันทึก
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}