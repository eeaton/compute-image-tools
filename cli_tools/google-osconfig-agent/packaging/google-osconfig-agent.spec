# Copyright 2018 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

Name: google-osconfig-agent
Version: %{_version}
Release: 1%{?dist}
Summary: Google Compute Engine guest environment.
License: ASL 2.0
Url: https://github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/google-osconfig-agent
Source0: %{name}_%{version}.orig.tar.gz

ExclusiveArch: %{go_arches}
%if 0%{?rhel} == 7
BuildRequires: systemd
%endif
BuildRequires: golang >= 1.11

%description
Contains the OSConfig agent binary and startup scripts

%prep
%autosetup

%build
GOPATH=%{gopath} go build -ldflags="-linkmode=external" -o google_osconfig_agent

%install
install -d %{buildroot}%{_bindir}
install -p -m 0755 google_osconfig_agent %{buildroot}%{_bindir}/google_osconfig_agent
%if 0%{?rhel} == 7
install -d %{buildroot}%{_unitdir}
install -p -m 0644 %{name}.service %{buildroot}%{_unitdir}
%else
install -d %{buildroot}/etc/init
install -p -m 0644 %{name}.conf %{buildroot}/etc/init
%endif

%files
%defattr(-,root,root,-)
%{_bindir}/google_osconfig_agent
%if 0%{?rhel} == 7
%{_unitdir}/%{name}.service
%else
/etc/init/%{name}.conf
%endif