import React, { useCallback, useContext } from "react";
import Box from "@mui/material/Box";
import IconButton from "@mui/material/IconButton";
import { AppContextProps } from "../../../types/declarations/app";
import { AppContext } from "../../../App";
import { useNavigate } from "react-router-dom";
import { getBaseHref } from "../../../utils";
import Chip from "@mui/material/Chip";

export default function AccountMenu() {
  const navigate = useNavigate();
  const { userInfo } = useContext<AppContextProps>(AppContext);

  const url =
    userInfo?.email === ""
      ? `${getBaseHref()}/auth/local/v1/logout`
      : `${getBaseHref()}/auth/v1/logout`;

  const handleLogout = useCallback(async () => {
    try {
      const response = await fetch(url);
      if (response.ok) {
        navigate("/login");
      }
      // TODO on failure?
    } catch (e: any) {
      // TODO on failure?
    }
  }, [url]);

  if (!userInfo) {
    // Non-auth, hide account menu
    return undefined;
  }

  return (
    <React.Fragment>
      <Box
        sx={{ display: "flex", alignItems: "center", textAlign: "center" }}
        data-testid="account-menu"
      >
        <IconButton onClick={handleLogout} size="small" sx={{ ml: 2 }}>
          <Chip
            label="Log out"
            sx={{ backgroundColor: "#00D0E0", color: "#000" }}
          />
        </IconButton>
      </Box>
    </React.Fragment>
  );
}
