import React, { useEffect, useState } from "react";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import { Link } from "react-router-dom";
import axios from "axios";

const Dashboard = () => {
  const [newsCount, setNewsCount] = useState(0);
  const [courseCount, setCourseCount] = useState(0);
  const [subjectCount, setSubjectCount] = useState(0);
  const [activityCount, setActivityCount] = useState(0);
  const [staffCount, setStaffCount] = useState(0);
  const [userCount, setUserCount] = useState(0);

  useEffect(() => {
    fetchDashboardData();
  }, []);

  const fetchDashboardData = async () => {
    try {
      const newsRes = await axios.get("http://localhost:8080/api/v1/admin/news");
      const courseRes = await axios.get("http://localhost:8080/api/v1/admin/course");
      const subjectRes = await axios.get("http://localhost:8080/api/v1/admin/subject");
    //   const activityRes = await axios.get("http://localhost:8080/api/v1/admin/activity");
      const staffRes = await axios.get("http://localhost:8080/api/v1/admin/personnel");
    //   const userRes = await axios.get("http://localhost:8080/api/v1/admin/users");

      setNewsCount(newsRes.data.length);
      setCourseCount(courseRes.data.length);
      setSubjectCount(subjectRes.data.length);
    //   setActivityCount(activityRes.data.length);
      setStaffCount(staffRes.data.length);
    //   setUserCount(userRes.data.length);
    } catch (error) {
      console.error("เกิดข้อผิดพลาดในการโหลดข้อมูล Dashboard:", error);
    }
  };

  return (
    <>
      <Headers />
      <Navbar />

      <div className="container text-center mt-4">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>

          <div className="col-sm-9">
            <h2 className="mb-4">แดชบอร์ดผู้ดูแลระบบ</h2>

            <div className="row mb-4">
              <div className="col-md-4">
                <div className="card text-white bg-primary mb-3">
                  <div className="card-body">
                    <h5 className="card-title">ข่าวสารทั้งหมด</h5>
                    <p className="card-text display-5">{newsCount}</p>
                  </div>
                </div>
              </div>

              <div className="col-md-4">
                <div className="card text-white bg-info mb-3">
                  <div className="card-body">
                    <h5 className="card-title">หลักสูตรทั้งหมด</h5>
                    <p className="card-text display-5">{courseCount}</p>
                  </div>
                </div>
              </div>

              <div className="col-md-4">
                <div className="card text-white bg-success mb-3">
                  <div className="card-body">
                    <h5 className="card-title">รายวิชา</h5>
                    <p className="card-text display-5">{subjectCount}</p>
                  </div>
                </div>
              </div>

              <div className="col-md-4">
                <div className="card text-white bg-dark mb-3">
                  <div className="card-body">
                    <h5 className="card-title">กิจกรรมทั้งหมด</h5>
                    <p className="card-text display-5">{activityCount}</p>
                  </div>
                </div>
              </div>

              <div className="col-md-4">
                <div className="card text-white bg-warning mb-3">
                  <div className="card-body">
                    <h5 className="card-title">บุคลากรทั้งหมด</h5>
                    <p className="card-text display-5">{staffCount}</p>
                  </div>
                </div>
              </div>

              <div className="col-md-4">
                <div className="card text-white bg-secondary mb-3">
                  <div className="card-body">
                    <h5 className="card-title">ผู้ใช้งาน</h5>
                    <p className="card-text display-5">{userCount}</p>
                  </div>
                </div>
              </div>
            </div>

            <div className="row">
              <div className="col">
                <Link to="/admin/news" className="btn btn-outline-info me-2 mb-2">
                  จัดการข่าวสาร
                </Link>
                <Link to="/admin/course" className="btn btn-outline-primary me-2 mb-2">
                  จัดการหลักสูตร
                </Link>
                <Link to="/admin/users" className="btn btn-outline-success me-2 mb-2">
                  จัดการผู้ใช้
                </Link>
                <Link to="/admin/activities" className="btn btn-outline-dark me-2 mb-2">
                  จัดการกิจกรรม
                </Link>
                <Link to="/admin/personnel" className="btn btn-outline-warning me-2 mb-2">
                  จัดการบุคลากร
                </Link>
                <Link to="/admin/subject" className="btn btn-outline-secondary mb-2">
                  จัดการรายวิชา
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
};

export default Dashboard;
