import React from "react";
import "./css/Login.css";
import { Divider } from "antd";
import { useState } from "react";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isHovered, setIsHovered] = React.useState(false);

  const handleMouseEnter = () => {
    setIsHovered(true);
  };

  const handleMouseLeave = () => {
    setIsHovered(false);
  };

  const handleLogin = async () => {
    // Tạo đối tượng formData chứa dữ liệu đăng nhập
    const formData = {
      username: username,
      password: password,
    };

    // Gửi đối tượng formData đến backend dưới dạng JSON
    const response = await fetch("http://localhost:3030/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    });

    // Xử lý phản hồi từ backend
    const data = await response.json();
    console.log(data);
  };

  return (
    <div className="login">
      <img
        src="https://getheavy.com/wp-content/uploads/2019/12/spotify2019-830x350.jpg"
        alt=""
      />
      <br />
      <p className="slogan-loginn">To continue, log in to Sportify.</p>
      <button
        className="btn-in-login"
        style={{
          marginTop: 0,
          backgroundColor: isHovered ? "white" : "rgb(38, 67, 255)",
          color: "black",
          padding: "10px 20px",
          border: "none",
          transition: "background-color 0.3s ease-in-out",
        }}
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}
      >
        CONTINUE WITH FACEBOOK
      </button>
      <button className="btn-in-login">CONTINUE WITH APPLE</button>
      <button type="submit" className="btn-in-login">
        CONTINUE WITH GOOGLE
      </button>
      <button type="submit" className="btn-in-login">
        CONTINUE WITH PHONE NUMBER
      </button>
      <>
        <Divider style={{ marginTop: 10, fontSize: 10 }}>OR</Divider>
      </>
      <div className="input-login" style={{ paddingTop: 0 }}>
        <label htmlFor="" className="label-input-loginn">
          Email address or username
        </label>
        <br />
        <input
          type="text"
          className="input-login-username"
          placeholder="Email address or username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <br />
        <label
          htmlFor=""
          className="label-input-loginn"
          style={{ marginRight: 230 }}
        >
          Password:
        </label>
        <input
          type="text"
          className="input-login-username"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <br />
        <button
          type="submit"
          className="btn-in-login-login"
          onClick={handleLogin}
        >
          LOG IN
        </button>
      </div>
    </div>
  );
}

export default Login;
