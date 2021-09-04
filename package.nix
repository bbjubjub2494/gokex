{ buildGoModule, lib }:

buildGoModule {
  name = "gokex";
  src = ./.;
  vendorSha256 = "sha256-ucXY/yplVut6wvVRProB4l1Hcx8dCym0EC1hgiCRTZ8=";

  meta = with lib; {
    description = "OKEx API client";
    license = licenses.lgpl3Plus;
    maintainers = with maintainers; [ lourkeur ];
  };
}
